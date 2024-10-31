package middlewares

import (
	"VincentLimarus/go-register-gmail-otps/configs"
	"VincentLimarus/go-register-gmail-otps/models/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func RequireUserAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			log.Fatal(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.SetCookie("Authorization", "", -1, "/", "", false, false) 
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			var user database.Users = database.Users{}

			db := configs.GetDB()
			_ = db.Where("id = ?", claims["sub"]).First(&user)

			if user.ID == uuid.Nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}