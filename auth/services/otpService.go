package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strive_go/config"

	"github.com/gin-gonic/gin"
)

func SendOtp(phone string, otp int, c *gin.Context) error {
	key := config.GetEnv("2FACTOR_KEY")
	if key == "" {
		return fmt.Errorf("Env 2FACTOR_KEY is missing, hence OTP not sent.")
	}
	api := fmt.Sprintf("https://2factor.in/API/V1/%s/SMS/+91%s/%d/OTP1", key, phone, otp)

	resp, err := http.Get(api)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//unmarshal Json
	var response map[string]interface{}
	err = json.Unmarshal(data, &response)

	if response["Status"] == "Success" {
		return nil
	}

	return fmt.Errorf("Unknown Error occured in SMS Service")
}
