package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strive_go/config"

	"github.com/razorpay/razorpay-go"
)

func CreateOrder(amount int, currency string) (string, int, string, error) {
	client := razorpay.NewClient(config.GetEnv("RAZORPAY_KEY_ID"), config.GetEnv("RAZORPAY_KEY_SECRET"))

	data := map[string]interface{}{
		"amount":   amount * 100, // amount in the smallest currency unit
		"currency": currency,
		"receipt":  "receipt_order_74394",
	}

	order, err := client.Order.Create(data, nil)
	if err != nil {
		return "", 0, "", err
	}

	return order["id"].(string), order["amount"].(int), order["currency"].(string), nil
}

func VerifyPayment(orderID, paymentID, signature string) (bool, error) {
	keySecret := config.GetEnv("RAZORPAY_KEY_SECRET")
	h := hmac.New(sha256.New, []byte(keySecret))
	h.Write([]byte(orderID + "|" + paymentID))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	return expectedSignature == signature, nil
}
