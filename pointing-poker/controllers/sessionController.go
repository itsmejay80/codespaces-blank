package controllers

import (
	"net/http"
	"pointing-poker/initializers"
	"pointing-poker/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {

	var body struct {
		SessionName string
		Username    string
	}

	c.BindJSON(&body)

	user := models.User{Username: body.Username}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	session := models.Session{CreatedAt: time.Now(), SessionName: body.SessionName, Owner: user}

	result = initializers.DB.Create(&session)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": session.SessionID})

}

func JoinSession(c *gin.Context) {

	var session models.Session
	result := initializers.DB.First(&session, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": session})
}
