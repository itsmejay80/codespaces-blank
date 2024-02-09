package main

import (
	"pointing-poker/initializers"

	models "pointing-poker/models"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Session{})
}
