package helpers

import (
	"errors"
	"fmt"
	"time"

	"VincentLimarus/go-register-gmail-otps/configs"
	"VincentLimarus/go-register-gmail-otps/models/database"
	"VincentLimarus/go-register-gmail-otps/models/outputs"
	"VincentLimarus/go-register-gmail-otps/models/requests"
	"VincentLimarus/go-register-gmail-otps/models/responses"
	"VincentLimarus/go-register-gmail-otps/utils"

	"gorm.io/gorm"
)

func UserRequestRegisterByEmail(RequestRegisterByEmailRequest requests.RequestRegisterByEmailRequest) (int, interface{}) {
	now := time.Now()

	otp, err := utils.GenerateOTP()
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	otpExpiry := now.Add(time.Minute * 5)

	var user database.Users
	var otps database.Otps

	db := configs.GetDB()

	err = db.Where("email = ?", RequestRegisterByEmailRequest.Email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.Email = RequestRegisterByEmailRequest.Email
			user.CreatedAt = now
			user.UpdatedAt = now
			user.IsActive = false

			err = db.Create(&user).Error
			if err != nil {
				output := outputs.InternalServerErrorOutput{
					Code:    500,
					Message: "Internal Server Error: Failed to create user: " + err.Error(),
				}
				return 500, output
			}
		} else {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error: " + err.Error(),
			}
			return 500, output
		}
	} else {
		err = db.Model(&user).Updates(database.Users{UpdatedAt: now, IsActive: false}).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error: Failed to update user: " + err.Error(),
			}
			return 500, output
		}
	}

	otps.OTP = otp
	otps.OTPExpiry = otpExpiry
	otps.UserID = user.ID
	otps.CreatedAt = now
	otps.UpdatedAt = now

	err = db.Create(&otps).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: Failed to create OTP: " + err.Error(),
		}
		return 500, output
	}

	if err = utils.SendOTPEmail(RequestRegisterByEmailRequest.Email, otp); err != nil {
		return fmt.Println("failed to send OTP email: %w", err)
	}

	output := outputs.RegisterByEmailOutput{
		Message: "Success",
		Code:    200,
		Email:   RequestRegisterByEmailRequest.Email,
		OTP:     otp,
	}
	return 200, output
}

func VerifyOTPByEmailRequest(VerifyOTPByEmailRequest requests.VerifyOTPByEmailRequest) (int, interface{}) {
	now := time.Now()

	var user database.Users
	var otps database.Otps

	db := configs.GetDB()

	err := db.Where("email = ?", VerifyOTPByEmailRequest.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			output := outputs.NotFoundOutput{
				Code:    404,
				Message: "Not Found: User with email " + VerifyOTPByEmailRequest.Email + " not found",
			}
			return 404, output
		}
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	err = db.Where("user_id = ? AND otp = ?", user.ID, VerifyOTPByEmailRequest.OTP).First(&otps).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			output := outputs.NotFoundOutput{
				Code:    404,
				Message: "Not Found: OTP with user id " + user.ID.String() + " and OTP " + VerifyOTPByEmailRequest.OTP + " not found",
			}
			return 404, output
		}
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	if otps.OTPExpiry.Before(now) {
		otps.IsActive = false
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: OTP with user id " + user.ID.String() + " and OTP " + VerifyOTPByEmailRequest.OTP + " has expired",
		}
		err = db.Delete(&otps).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error: " + err.Error(),
			}
			return 500, output
		}
		return 400, output
	}

	otps.IsActive = true
	otps.UpdatedAt = now
	err = db.Save(&otps).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	err = db.Model(&user).Updates(database.Users{IsActive: true, UpdatedAt: now}).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: Failed to update user: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.VerifyOTPByEmailOutput{
		Message: "Success",
		Code:    200,
		Data: responses.VerifyOTPResponse{
			Message: "OTP with user id " + user.ID.String() + " and OTP " + VerifyOTPByEmailRequest.OTP + " is valid",
			Email:   user.Email,
			IsValid: true,
		},
	}

	err = db.Delete(&otps).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	return 200, output
}
