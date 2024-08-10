package dice

import (
	"fmt"
	"strconv"
	gameplay "strive_go/games/services/gamePlay"
)

func DiceOutput(rollOver string, amount string) (int, error) {

	// convert string to int
	amt, err := strconv.Atoi(amount)
	if err != nil {
		return -1, fmt.Errorf("gamePlay/dice:\n error converting amount to int:\n %v", err)
	}
	roll, _ := strconv.Atoi(rollOver)

	externamWinchance := 100 - roll

	multiplier := 99 / externamWinchance

	if gameplay.ComputeResult(float64(multiplier), float64(amt), "") {
		// handle game win
		// TODO: return random number between rollOver and 100
		return 0, nil
	}
	// handle game lose
	// TODO: return random number between 0 and rollOver
	return 0, nil
}
