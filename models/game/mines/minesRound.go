package mines

import "gorm.io/gorm"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}
type MinesRound struct {
	gorm.Model

	PayOutMultiplier float64   `json:"payOutMultiplier"`
	AmountMultiplier float64   `json:"amountMultiplier"`
	Amount           float64   `json:"amount"`
	Payout           float64   `json:"payout"`
	UserID           uint64    `json:"userID"`
	User             User      `json:"user"`
	BetNo            int64     `json:"betNo"`
	BoxNo            int64     `json:"boxNo"`
	MinesDataID      uint      `json:"minesDataID" gorm:"index"`
	MinesData        MinesData `json:"minesData"`
}
