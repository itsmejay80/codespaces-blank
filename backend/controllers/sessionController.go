package controllers

import (
	"net/http"
	"pointing-poker/initializers"
	"pointing-poker/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSession(c *gin.Context) {

	var body struct {
		SessionName string
		Username    string
	}

	c.BindJSON(&body)

	user := models.User{Username: body.Username}

	user_result := initializers.DB.Create(&user)

	if user_result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	// check if a session with same name exists

	var existingSession models.Session
	if err := initializers.DB.Where("session_name = ?", body.SessionName).First(&existingSession).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session with this name already exists"})
		return
	}

	session := models.Session{CreatedAt: time.Now(), SessionName: body.SessionName, OwnerID: user.UserID}

	result := initializers.DB.Create(&session)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": session.SessionID})

}

func JoinSession(c *gin.Context) {

	var body struct {
		Username string
	}

	c.BindJSON(&body)

	var session models.Session
	result := initializers.DB.Preload("Users").Preload("Owner").First(&session, c.Param("id"))

	user := models.User{Username: body.Username}
	initializers.DB.Create(&user)

	initializers.DB.Model(&session).Association("Users").Append(&user)

	initializers.DB.Preload("Users").Preload("Owner").First(&session, session.SessionID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": session})
}

func LeaveSession(c *gin.Context) {

	var body struct {
		Username string
	}

	c.BindJSON(&body)

	var session models.Session
	result := initializers.DB.Preload("Users").Preload("Owner").First(&session, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find session"})
		return
	}

	// Find the user to remove from the session
	var user models.User
	result = initializers.DB.Where("username = ?", body.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find user"})
		return
	}

	// Remove the user from the session
	initializers.DB.Model(&session).Association("Users").Delete(&user)

	// Delete the user from the database
	initializers.DB.Delete(&user)

	// Reload the session to get the updated Users slice
	initializers.DB.Preload("Users").Preload("Owner").First(&session, session.SessionID)

	c.JSON(http.StatusOK, gin.H{"session": session})

}
