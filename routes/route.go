package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"VincentLimarus/go-register-gmail-otps/controllers/services"
)

func SetupRoutes() *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(config))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "go-register-gmail-otps Endpoints")
	})

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "go-register-gmail-otps API Version 1.0.0")
	})

	base := router.Group("/api/v1")
	services.AccountService(base)

	return router
}
