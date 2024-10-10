package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"net/url"
	"os"
)

func GenerateRandomState(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	state := make([]byte, length)
	for i := range state {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		state[i] = charset[num.Int64()]
	}
	return string(state), nil
}

func GenerateGoogleLoginURL(state string) string {
	clientID := os.Getenv("GOOGLE_CLIENT_ID") 
	redirectURI := os.Getenv("GOOGLE_REDIRECT_URI") 

	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("response_type", "code")
	params.Set("scope", "email profile")
	params.Set("state", state)

	return fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?%s", params.Encode())
}

func GenerateOTP() (string, error) {
	otp := ""

	firstDigit, err := rand.Int(rand.Reader, big.NewInt(9)) 
	if err != nil {
		return "", err
	}
	otp += fmt.Sprintf("%d", firstDigit)

	for i := 1; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10)) 
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", num) 
	}

	return otp, nil
}

func SendOTPEmail(userEmail, otp string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD") 

	to := []string{userEmail}
	subject := "Hello There!, Here is your OTP code"
	body := "Your OTP code is: " + otp

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	msg := []byte("From: " + from + "\n" +
		"To: " + userEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	return err
}
