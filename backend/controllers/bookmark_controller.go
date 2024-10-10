package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/keoy7am/GoNav/backend/database"
	"github.com/keoy7am/GoNav/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateBookmark(c *gin.Context) {
	var bookmark models.Bookmark
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert \n back to actual newline characters
	bookmark.Description = strings.ReplaceAll(bookmark.Description, "\\n", "\n")

	// Validate URL format
	urlPattern := regexp.MustCompile(`^https?:\/\/`)
	if !urlPattern.MatchString(bookmark.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format. Must start with http:// or https://"})
		return
	}

	username := c.GetString("username")
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user information"})
		return
	}
	bookmark.UserID = user.ID

	// Validate if the category exists and belongs to the user
	if bookmark.CategoryID != 0 {
		var category models.Category
		if err := database.DB.Where("id = ? AND user_id = ?", bookmark.CategoryID, user.ID).First(&category).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
	}

	if err := database.DB.Create(&bookmark).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create bookmark"})
		return
	}

	c.JSON(http.StatusCreated, bookmark)
}

func GetBookmarks(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user information"})
		return
	}

	var bookmarks []models.Bookmark
	if err := database.DB.Where("user_id = ?", user.ID).Find(&bookmarks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get bookmarks"})
		return
	}

	for i := range bookmarks {
		// Convert actual newline characters to \n
		bookmarks[i].Description = strings.ReplaceAll(bookmarks[i].Description, "\n", "\\n")
	}

	c.JSON(http.StatusOK, bookmarks)
}

func UpdateBookmark(c *gin.Context) {
	id := c.Param("id")
	var bookmark models.Bookmark
	if err := database.DB.First(&bookmark, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}

	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert \n back to actual newline characters
	bookmark.Description = strings.ReplaceAll(bookmark.Description, "\\n", "\n")

	// Validate URL format
	urlPattern := regexp.MustCompile(`^https?:\/\/`)
	if !urlPattern.MatchString(bookmark.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format. Must start with http:// or https://"})
		return
	}

	if err := database.DB.Save(&bookmark).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update bookmark"})
		return
	}

	c.JSON(http.StatusOK, bookmark)
}

func DeleteBookmark(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DB.Delete(&models.Bookmark{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete bookmark"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bookmark deleted"})
}

func GetPublicBookmarks(c *gin.Context) {
	publicId := c.Param("publicId")

	var settings models.UserSettings
	if err := database.DB.Where("public_id = ? AND enable_public_bookmarks = ?", publicId, true).First(&settings).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Public bookmarks not found"})
		return
	}

	if settings.EnableAccessCode {
		accessCode := c.Query("access_code")
		if accessCode != settings.AccessCode {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access code"})
			return
		}
	}

	var bookmarks []models.Bookmark
	if err := database.DB.Where("user_id = ? AND is_private = ?", settings.UserID, false).Find(&bookmarks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get public bookmarks"})
		return
	}

	var categories []models.Category
	if err := database.DB.Where("user_id = ? AND is_private = ?", settings.UserID, false).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get public categories"})
		return
	}

	// Handle newline characters
	for i := range bookmarks {
		bookmarks[i].Description = strings.ReplaceAll(bookmarks[i].Description, "\n", "\\n")
	}

	response := gin.H{
		"bookmarks":  bookmarks,
		"categories": categories,
	}

	c.JSON(http.StatusOK, response)
}
