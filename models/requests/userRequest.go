package requests

type RequestRegisterByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPByEmailRequest struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}
