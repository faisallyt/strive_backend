package models

import (
	"gorm.io/gorm"
)

// Bet struct
type Bet struct {
	gorm.Model
	ID               uint    `json:"id" gorm:"primaryKey autoIncrement:true"`
	GameID           uint    `json:"game_id"`
	Username         string  `json:"username"`
	BetAmount        float64 `json:"amount"`
	InputMultiplier  float64 `json:"user_multiplier"`
	OutputMultiplier float64 `json:"output_multiplier"`
	WinChance        float64 `json:"win_chance"`
	WinAmount        float64 `json:"win_amount"`
	IsWin            bool    `json:"is_win"`
	Date             string  `json:"date"`
	Time             string  `json:"time"`
	Gamedata         uint    `gorm:"foreignKey:GameID"`
}
