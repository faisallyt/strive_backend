package functions

import (
	"fmt"
	"strive_go/db"
	"strive_go/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InserUserAuth(userAuth models.UserAuth) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userAuth.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	userAuth.Password = string(hashedPassword)

	//inserting user in  the database

	result := db.Instance.Create(&userAuth)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateOtp(userAuth models.UserAuth, otp uint) error {
	var existingUser models.UserAuth

	result := db.Instance.First(&existingUser, userAuth.ID)

	if result.Error != nil {
		return result.Error
	}

	existingUser.OTP = otp

	updateResult := db.Instance.Save(&existingUser)

	if updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}

func ChangeUserStatus(userAuth models.UserAuth) error {
	var existingUser models.UserAuth

	result := db.Instance.First(&existingUser, userAuth.ID)

	if result.Error != nil {
		return result.Error
	}

	existingUser.Verified = true
	updateResult := db.Instance.Save(&existingUser)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}

func UpdateRefreshToken(id uint, newRefreshToken string) error {
	var existingUser models.UserAuth

	result := db.Instance.First(&existingUser, id)

	if result.Error != nil {
		return result.Error
	}

	existingUser.RefreshToken = newRefreshToken
	// existingUser.UserID = nil
	updateResult := db.Instance.Save(&existingUser)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}

func ValidateLogin(email, username, phone, password string) error {
	var userAuth models.UserAuth

	//check if email exists
	if email != "" {
		result := db.Instance.Where("email =?", email).First(&userAuth)
		if result.Error != nil {
			return result.Error
		}
	}

	//check if username exists
	if username != "" {
		result := db.Instance.Where("username =?", username).First(&userAuth)
		if result.Error != nil {
			return result.Error
		}
	}

	//check if phone exists
	if phone != "" {
		result := db.Instance.Where("phone =?", phone).First(&userAuth)
		if result.Error != nil {
			return result.Error
		}
	}

	//check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func UserExists(field, value string) (bool, error) {
	var userAuth models.UserAuth
	result := db.Instance.Where(field+"=?", value).First(&userAuth)
	if result.Error == gorm.ErrRecordNotFound {
		return false, result.Error
	} else if result.Error != nil {
		fmt.Println(result.Error)
		return false, result.Error
	}
	// Check if user exists and is verified
	userExists := result.Error == nil && userAuth.Verified
	return userExists, nil
}

func DeleteUserAuth(userAuth models.UserAuth) error {
	result := db.Instance.Delete(&userAuth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ChangeUsernameInDB(currentUsername, newUsername string) error {
	var userAuth models.UserAuth

	result := db.Instance.Where("username=?", currentUsername).First(&userAuth)

	if result.Error != nil {
		return result.Error
	}
	userAuth.Username = newUsername
	result = db.Instance.Save(&userAuth)

	return nil
}

func GetUserAuthFromID(id uint) (models.UserAuth, error) {
	var userAuth models.UserAuth
	result := db.Instance.Where("id=?", id).First(&userAuth)
	if result.Error != nil {
		return userAuth, result.Error
	}
	return userAuth, nil
}
