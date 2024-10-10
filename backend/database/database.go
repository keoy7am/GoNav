package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/keoy7am/GoNav/backend/models"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbType := viper.GetString("database.type")
	dbName := viper.GetString("database.name")
	dbPath := viper.GetString("database.path")

	fullPath := filepath.Join(dbPath, dbName)
	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create database directory: %v", err)
	}

	if dbType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	} else {
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自動遷移
	err = DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Bookmark{}, &models.UserSettings{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	fmt.Println("Database connected and migrated successfully")
}
