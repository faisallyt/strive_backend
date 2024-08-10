package models

// UserInfo struct combines information from User and UserAuth
type UserInfo struct {
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Country      string  `json:"country"`
	DateOfBirth  string  `json:"date_of_birth"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	PostalCode   string  `json:"postal_code"`
	Occupation   string  `json:"occupation"`
	AvatarID     string  `json:"avatar_id"`
	IsActive     bool    `json:"is_active"`
	RefreshToken string  `json:"refresh_token"`
	Balance      float64 `json:"balance"`
	Aadhar       string  `json:"aadhar"`
	Verified     bool    `json:"verified"`
}
