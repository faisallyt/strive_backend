package models

import (
	"gorm.io/gorm"
)

// User struct

type User struct {
	gorm.Model
	FirstName   string  `json:"name"`
	LastName    string  `json:"lastname"`
	Country     string  `json:"country"`
	DateOfBirth string  `json:"dob"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	PostalCode  string  `json:"pincode"`
	Occupation  string  `json:"ocupation"`
	AvatarID    string  `json:"avatar"`
	IsActive    bool    `json:"is_active"`
	Balance     float64 `json:"balance"`
	Aadhar      string  `json:"aadhar"`
	UserAuthID  uint    `json:"user_auth_id" gorm:"index"`
	UserAuth    UserAuth
}
