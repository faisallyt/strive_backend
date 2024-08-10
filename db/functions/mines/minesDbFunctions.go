package mines

import (
	"errors"
	"math/rand"
	"strive_go/db"
	"strive_go/models"
	"strive_go/models/game/mines"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func CheckActiveGame(username string) (bool, *uint, error) {
	var games []mines.MinesData
	err := db.Instance.Where("username=? AND is_active =?", username, true).Find(&games).Error

	if err != nil {
		return false, nil, err
	}
	count := len(games)

	if count == 0 {
		return false, nil, nil
	} else if count == 1 {
		return true, &games[0].ID, nil
	} else {
		return true, nil, errors.New("multiple active games found for user")
	}

}

func StartNewGame(username string, amount float64, minesCount int64) error {
	newGame := mines.MinesData{
		Amount:         amount,
		MinesCount:     minesCount,
		IsActive:       true,
		ChosenCount:    0,
		Username:       username,
		WinningTillNow: amount,
	}

	result := db.Instance.Create(&newGame)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FetchCurrentGame(gameID *uint) (*mines.CurrentGameData, error) {
	if gameID == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var minesData mines.MinesData

	err := db.Instance.First(&minesData, *gameID).Error
	if err != nil {
		return nil, err
	}

	//fetch  the related MinesRound  record

	var minesRound []mines.MinesRound
	err = db.Instance.Where("mines_data_id =?", *gameID).Find(&minesRound).Error
	if err != nil {
		return nil, err
	}

	//Map the MinesRound data to the output struct

	rounds := make([]mines.MinesRoundData, len(minesRound))

	for i, round := range minesRound {
		rounds[i] = mines.MinesRoundData{
			Amount:           round.Amount,
			BoxNo:            round.BoxNo,
			BetNo:            round.BetNo,
			Payout:           round.Payout,
			AmountMultiplier: round.AmountMultiplier,
			PayOutMultiplier: round.PayOutMultiplier,
		}
	}

	currentGameData := &mines.CurrentGameData{
		Amount:     minesData.Amount,
		MinesCount: minesData.MinesCount,
		Rounds:     rounds,
	}

	return currentGameData, nil

}
func FindCurrentMinesData(gameID *uint) (*mines.MinesData, error) {
	if gameID == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var minesData mines.MinesData

	err := db.Instance.First(&minesData, *gameID).Error
	if err != nil {
		return nil, err
	}
	return &minesData, nil
}

func IsBoxChosenAlready(chosenBoxes pq.Int64Array, boxNO int64) bool {
	for _, chosenBox := range chosenBoxes {
		if chosenBox == boxNO {
			return true
		}
	}
	return false
}

func DebitUserBalance(username string, activeGame *mines.MinesData) error {
	var user *models.User
	err := db.Instance.Where("username=?", username).First(&user).Error

	if err != nil {
		return err
	}
	user.Balance -= activeGame.Amount
	result := db.Instance.Save(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil

}

func CreditUserBalance(multiplier float64, activeGame *mines.MinesData) error {
	activeGame.WinningTillNow = activeGame.Amount * multiplier
	result := db.Instance.Save(&activeGame)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddBoxesWithMines(activeGame *mines.MinesData, boxNo, minesCount int64) error {
	if activeGame == nil {
		return errors.New("active game can not be nil")
	}

	//Ensure if minesCount is Valid

	if minesCount <= 0 || minesCount >= 25 {
		return errors.New("invalid mines count")
	}

	//create map of chosen Boxes for quick lookup

	chosenBoxesMap := make(map[int64]bool)

	for _, box := range activeGame.UserChosenBoxes {
		chosenBoxesMap[box] = true
	}

	//create an array to hold the boxes with mines
	var boxesWithMines []int64

	if chosenBoxesMap[boxNo] {
		return errors.New("box already chosen and has true value ")
	}

	boxesWithMines = append(boxesWithMines, boxNo)
	chosenBoxesMap[boxNo] = true

	//generate unique boxes with mines

	rand.Seed(time.Now().UnixNano())
	for len(boxesWithMines) < int(minesCount) {
		newBox := rand.Int63n(25)
		if !chosenBoxesMap[newBox] {
			boxesWithMines = append(boxesWithMines, newBox)
			chosenBoxesMap[newBox] = true
		}

	}
	activeGame.BoxesWithMines = boxesWithMines
	result := db.Instance.Save(activeGame)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func EndMinesGame(activeGame *mines.MinesData) error {
	activeGame.IsActive = false
	result := db.Instance.Save(activeGame)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddNewRoundsInDB(activeGame *mines.MinesData, boxNo int64, multiplier float64) error {
	round := mines.MinesRound{
		MinesDataID:      activeGame.ID,
		BoxNo:            boxNo,
		BetNo:            activeGame.ChosenCount,
		Amount:           activeGame.Amount,
		Payout:           activeGame.WinningTillNow,
		AmountMultiplier: multiplier,
		PayOutMultiplier: multiplier,
	}

	result := db.Instance.Create(&round)
	if result.Error != nil {
		return result.Error
	}

	userChosenBoxes := append(activeGame.UserChosenBoxes, boxNo)

	activeGame.UserChosenBoxes = userChosenBoxes
	activeGame.ChosenCount++
	result = db.Instance.Save(activeGame)

	if result.Error != nil {
		return result.Error
	}

	return nil
	// update the winningTillNow in the MinesData table
}
