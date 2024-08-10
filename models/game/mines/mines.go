package mines

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Mines struct
type MinesData struct {
	gorm.Model
	Amount          float64       `json:"amount"`
	MinesCount      int64         `json:"mines_count"`
	IsActive        bool          `json:"is_active"`
	Username        string        `json:"user_id"`
	BoxesWithResult pq.Int64Array `json:"boxes_with_result" gorm:"type:integer[]"`
	UserChosenBoxes pq.Int64Array `json:"user_chosen_boxes" gorm:"type:integer[]"`
	BoxesWithMines  pq.Int64Array `json:"boxes_with_mines" gorm:"type:integer[]"`
	WinningTillNow  float64       `json:"winning_till_now"`
	ChosenCount     int64         `json:"chosen_count"`
	MinesRounds     pq.Int64Array `json:"mines_rounds" gorm:"type:integer[]"`
}
