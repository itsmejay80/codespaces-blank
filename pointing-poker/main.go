package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pointing-poker/initializers"

	"pointing-poker/controllers"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Pointing Poker Server Running!!")
	})

	router.POST("/create-session", controllers.CreateSession)

	router.GET("/join-session/:id", controllers.JoinSession)

	router.Run(":8000")
}
