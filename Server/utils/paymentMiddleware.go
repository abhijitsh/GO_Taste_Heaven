package payment

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"log"
	"github.com/razorpay/razorpay-go"
)

// initiation of razorpay req
func Executerazorpay(price int, orderId string) (string, error) {

	client := razorpay.NewClient("YOUR_KEY_ID", "YOUR_SECRET")

	data := map[string]interface{}{
		"amount":   int(price) * 100,
		"currency": "INR",
		"receipt":  orderId,
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", errors.New("payment not initiated")
	}
	razorId, _ := body["id"].(string)
	return razorId, nil
}

// payment verification
func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := "YOUR_SECRET"
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("payment failed")
	} else {
		return nil
	}
}
