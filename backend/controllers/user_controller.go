package controllers

import (
	"net/http"
	"time"

	"github.com/keoy7am/GoNav/backend/database"
	"github.com/keoy7am/GoNav/backend/models"
	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to encrypt password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", loginInfo.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Read JWT key and expiration time here
	jwtKey := []byte(viper.GetString("jwt.secret"))
	expirationHours := viper.GetInt("jwt.expirationHours")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetUserInfo(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"isPublic": user.IsPublic,
	})
}

func UpdateUserSettings(c *gin.Context) {
	username := c.GetString("username")
	var settings struct {
		IsPublic        *bool   `json:"isPublic"`
		BasicAuthUser   *string `json:"basicAuthUser"`
		BasicAuthPasswd *string `json:"basicAuthPasswd"`
	}

	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if settings.IsPublic != nil {
		updates["is_public"] = *settings.IsPublic
	}
	if settings.BasicAuthUser != nil {
		updates["basic_auth_user"] = *settings.BasicAuthUser
	}
	if settings.BasicAuthPasswd != nil {
		updates["basic_auth_passwd"] = *settings.BasicAuthPasswd
	}

	if err := database.DB.Model(&models.User{}).Where("username = ?", username).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User settings updated"})
}
