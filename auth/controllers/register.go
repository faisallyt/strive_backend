package controllers

import (
	"math/rand/v2"
	"net/http"
	"strive_go/auth/services"
	"strive_go/auth/utils"
	"strive_go/auth/utils/validators"
	"strive_go/db/functions"
	"strive_go/models"

	"github.com/gin-gonic/gin"
)

// RegisterController Return the response of userData provided by client.
// RegisterController                godoc
// @title		 RegisterController
// @Summary      request for registeration/signUp
// @Description  Returns the response of data provided by client for the registration at our portal.
// @Tags         Auth
// @Produce      json
// @Param         email      formData  string  true  "Email address"
// @Param         username   formData  string  true  "Username"
// @Param         dob        formData  string  true  "Date of birth in YYYY-MM-DD format"
// @Param         password formData string true "Password"
// @Param         phone      formData  string  true  "Phone number with country code"
// @Success       200        {object}  map[string]interface{}  "ok"
// @Router       /api/v1/auth/user [put]
func RegisterController(c *gin.Context) {

	var registrationData struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Phone    string `form:"phone" binding:"required"`
		Password string `form:"password" binding:"required"`
		DOB      string `form:"dob" binding:"required"`
	}

	if err := c.ShouldBind(&registrationData); err != nil {
		utils.SendApiError(c, http.StatusBadGateway, "Invalid request data/format.")
		return
	}

	// Check for existing user with the same username

	if res, err := functions.UserExists("username", registrationData.Username); res == true && err == nil {
		utils.SendApiError(c, http.StatusConflict, "Username already exists")
		return
	}

	// Check for existing user with the same email
	if res, err := functions.UserExists("email", registrationData.Email); res == true && err == nil {
		utils.SendApiError(c, http.StatusConflict, "Email already exists")
		return
	}

	// Check for existing user with the same phone
	if res, err := functions.UserExists("phone", registrationData.Phone); res == true && err == nil {
		utils.SendApiError(c, http.StatusConflict, "Phone already exists")
		return
	}

	// Check for valid date of birth
	isvalidDOB, message := validators.IsvalidDOB(registrationData.DOB)

	if !isvalidDOB {
		utils.SendApiError(c, http.StatusBadRequest, message.Error())
		return
	}

	// validate phone number

	isValidPhone := validators.IsValidPhone(registrationData.Phone)

	if !isValidPhone {
		utils.SendApiError(c, http.StatusBadRequest, "Please follow standard phone number format.")
		return
	}

	//  validatev password
	isValidPassword := validators.IsValidPassword(registrationData.Password)
	if !isValidPassword {
		utils.SendApiError(c, http.StatusBadRequest, "Please follow standard password format.")
		return
	}

	// validate username
	IsValidUsername := validators.IsValidUsername(registrationData.Username)

	if !IsValidUsername {
		utils.SendApiError(c, http.StatusBadRequest, "Please follow standard username format.")
		return
	}

	// validate email
	isValidEmail, err := validators.IsValidEmail(registrationData.Email)
	if !isValidEmail {
		utils.SendApiError(c, http.StatusBadRequest, err.Error())
		return
	}

	userAuth := models.UserAuth{
		Username: registrationData.Username,
		Email:    registrationData.Email,
		Phone:    registrationData.Phone,
		Password: registrationData.Password,
		DOB:      registrationData.DOB,
		Verified: false,
	}

	//Generate a random 6 digit otp

	otp := rand.IntN(100000) + (rand.IntN(9)+1)*100000
	userAuth.OTP = uint(otp)

	//Insert the UserAuth in the database

	err = functions.InserUserAuth(userAuth)
	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Could not Register User Auth "+err.Error())
		return
	}

	err = services.SendOtp(registrationData.Phone, otp, c)
	if err != nil {
		functions.DeleteUserAuth(userAuth)
		utils.SendApiError(c, http.StatusInternalServerError, "Could not send OTP: "+err.Error())
		return
	}

	//send success message
	response := gin.H{
		"message":  "user registered partially ,Kindly verify phone number",
		"username": registrationData.Username,
	}

	utils.SendSuccessResponse(c, http.StatusOK, "registered partially", response)
	return
}
