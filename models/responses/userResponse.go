package responses

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        			uuid.UUID `json:"id"`
	Name      			string    `json:"name"`
	Email    		 	string    `json:"email"`
	ProfilePicture 		string `json:"profile_picture"`
	IsActive  			bool      `json:"is_active"`
	CreatedBy 			string    `json:"created_by"`
	UpdatedBy 			string    `json:"updated_by"`
	CreatedAt 			time.Time `json:"created_at"`
	UpdatedAt 			time.Time `json:"updated_at"`

	Otps 				[]OtpResponse `json:"otps"`
}

type RegisterWithEmailResponse struct {
	Message string `json:"message"`
	Email  	string `json:"email"`
	OTP   	string `json:"otp"`
}

type VerifyOTPResponse struct {
	Message string `json:"message"`
	Email  	string `json:"email"`
	IsValid bool `json:"is_valid"`
}