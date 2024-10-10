package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keoy7am/GoNav/backend/database"
	"github.com/keoy7am/GoNav/backend/models"
)

func GetSettings(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user information"})
		return
	}

	var settings models.UserSettings
	if err := database.DB.Where("user_id = ?", user.ID).First(&settings).Error; err != nil {
		if err.Error() == "record not found" {
			settings = models.UserSettings{UserID: user.ID}
			database.DB.Create(&settings)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get settings"})
			return
		}
	}

	c.JSON(http.StatusOK, settings)
}

func UpdateSettings(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user information"})
		return
	}

	var settings models.UserSettings
	if err := database.DB.Where("user_id = ?", user.ID).First(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get settings"})
		return
	}

	var newSettings models.UserSettings
	if err := c.ShouldBindJSON(&newSettings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update settings
	settings.EnablePublicBookmarks = newSettings.EnablePublicBookmarks
	settings.EnableAccessCode = newSettings.EnableAccessCode
	settings.AccessCode = newSettings.AccessCode

	if settings.EnablePublicBookmarks && settings.PublicID == "" {
		publicID, err := generatePublicID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate public ID"})
			return
		}
		settings.PublicID = publicID
	}

	// Ensure access code is set when enabled
	if settings.EnableAccessCode && settings.AccessCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access code must be set when enabled"})
		return
	}

	if err := database.DB.Save(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update settings"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func generatePublicID() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
