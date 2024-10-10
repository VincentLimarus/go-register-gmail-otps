package responses

import (
	"time"

	"github.com/google/uuid"
)

type OtpResponse struct {
	ID        uuid.UUID 	`json:"id"`
	UserID    uuid.UUID 	`json:"user_id"`
	OTP       string    	`json:"otp"`
	OTPExpiry string    	`json:"otp_expiry"`
	IsActive  bool      	`json:"is_active"`
	CreatedBy string    	`json:"created_by"`
	UpdatedBy string    	`json:"updated_by"`
	CreatedAt time.Time    	`json:"created_at"`
	UpdatedAt time.Time    	`json:"updated_at"`
}