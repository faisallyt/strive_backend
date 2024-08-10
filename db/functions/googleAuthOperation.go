package functions

import (
	"errors"
	"strive_go/db"
	"strive_go/models"

	"gorm.io/gorm"
)

func InsertUserByGoogleAuth(username string, email, refreshToken string) (models.UserAuth, error) {
	var user models.UserAuth

	if err := db.Instance.Where("email=?", email).First(&user).Error; err == nil {
		//User alread exists
		return models.UserAuth{}, errors.New("User already exists")
	} else if err := db.Instance.Where("username=?", username).First(&user).Error; err == nil {
		//User alread exists
		return models.UserAuth{}, errors.New("User already exists")
	} else if err != gorm.ErrRecordNotFound {
		return models.UserAuth{}, err
	}

	newUser := models.UserAuth{
		Username:     username,
		Email:        email,
		RefreshToken: refreshToken,
		Verified:     true,
	}

	if err := db.Instance.Save(&newUser).Error; err != nil {
		return models.UserAuth{}, err
	}

	return newUser, nil
}
