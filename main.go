package main

import (
	"VincentLimarus/go-register-gmail-otps/configs"
	"VincentLimarus/go-register-gmail-otps/models/database"
	"VincentLimarus/go-register-gmail-otps/routes"
)

func init() {
	configs.LoadEnviromentVar()
	configs.ConnectToDB()
}

func main() {
	db := configs.GetDB()
	db.AutoMigrate(&database.Users{}, &database.Otps{})

	router := routes.SetupRoutes()
	router.Run(":3000")
}