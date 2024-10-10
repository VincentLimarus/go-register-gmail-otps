package services

import (
	"fmt"
	"net/http"

	"VincentLimarus/go-register-gmail-otps/controllers/helpers"
	"VincentLimarus/go-register-gmail-otps/models/outputs"
	"VincentLimarus/go-register-gmail-otps/models/requests"

	"github.com/gin-gonic/gin"
)

func RegisterAccount(c *gin.Context) {
	var RegisterAccountRequestDTO requests.RequestRegisterByEmailRequest
	if err := c.ShouldBindJSON(&RegisterAccountRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)

		return
	}
	code, output := helpers.UserRequestRegisterByEmail(RegisterAccountRequestDTO)
	c.JSON(code, output)
}

func VerifyOTPByEmailRequest(c *gin.Context) {
	var VerifyOTPByEmailRequestDTO requests.VerifyOTPByEmailRequest
	if err := c.ShouldBindJSON(&VerifyOTPByEmailRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.VerifyOTPByEmailRequest(VerifyOTPByEmailRequestDTO)
	c.JSON(code, output)
}

func AccountService(router *gin.RouterGroup) {
	router.POST("/auth/register", RegisterAccount)
	router.POST("/auth/register/verify", VerifyOTPByEmailRequest)
}
