package models

import "gorm.io/gorm"

// Enums for different types of games

type GameType int

const (
	// Dice game
	Dice   GameType = 1
	Crash  GameType = 2
	Mines  GameType = 3
	Bowl   GameType = 4
	Hilo   GameType = 5
	Lucky7 GameType = 6
	Limbo  GameType = 7
)

// Game struct
type Game struct {
	gorm.Model
	GameType    GameType `json:"game_type"`
	GameName    string   `json:"game_name"`
	Description string   `json:"description"`
	IsActive    bool     `json:"is_active"`
	ActiveUsers int      `json:"active_users"`
}
