package striveauth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(id uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := RefreshTokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTsecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
