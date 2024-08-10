package functions

import (
	"strive_go/db"
	"strive_go/models"
)

func InsertUserInfo(user models.User) error {

	// insert User information into Db
	result := db.Instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	// return id of the user inserted

	return nil
}

func DeleteUser(user models.User) error {
	var existingUser models.User

	result := db.Instance.First(&existingUser, user.ID)

	if result.Error != nil {
		return result.Error
	}

	existingUser.IsActive = false

	updateResult := db.Instance.Save(&existingUser)

	if updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}

func GetUserInfo(username string) (models.UserInfo, error) {
	var userAuth models.UserAuth
	var user models.User
	var userInfo models.UserInfo

	//find the userAuth by username

	result := db.Instance.Where("username=?", username).First(&userAuth)
	if result.Error != nil {
		return userInfo, result.Error
	}

	//find the corresponding user by userId

	result = db.Instance.Where("user_auth_id=?", userAuth.ID).First(&user)
	if result.Error != nil {
		return userInfo, result.Error
	}

	//combine the Information  into UserInfo struct

	userInfo = models.UserInfo{
		Username:     userAuth.Username,
		Email:        userAuth.Email,
		Phone:        userAuth.Phone,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Country:      user.Country,
		DateOfBirth:  user.DateOfBirth,
		Address:      user.Address,
		City:         user.City,
		PostalCode:   user.PostalCode,
		Occupation:   user.Occupation,
		AvatarID:     user.AvatarID,
		IsActive:     user.IsActive,
		RefreshToken: userAuth.RefreshToken,
		Aadhar:       user.Aadhar,
		Verified:     userAuth.Verified,
	}
	return userInfo, nil
}
