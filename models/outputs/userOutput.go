package outputs

import (
	"VincentLimarus/go-register-gmail-otps/models/responses"
)

type RegisterByEmailOutput struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Email   string `json:"email"`
	OTP     string `json:"otp"`
}

type VerifyOTPByEmailOutput struct {
	Message string                      `json:"message"`
	Code    int                         `json:"code"`
	Data    responses.VerifyOTPResponse `json:"data"`
}

type VerifyOTPOutput struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Email   string `json:"email"`
	IsValid bool   `json:"is_valid"`
}
