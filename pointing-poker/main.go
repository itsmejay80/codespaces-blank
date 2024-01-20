package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pointing-poker/initializers"

	models "pointing-poker/models"

	"time"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Pointing Poker Server Running!!")
	})

	router.POST("/create-session", func(c *gin.Context) {
		session := models.Session{CreatedAt: time.Now()}

		result := initializers.DB.Create(&session)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"session": session.ID})
	})

	router.Run(":8000")
}
