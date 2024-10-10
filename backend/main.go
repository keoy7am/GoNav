package main

import (
	"log"

	"github.com/keoy7am/GoNav/backend/controllers"
	"github.com/keoy7am/GoNav/backend/database"
	"github.com/keoy7am/GoNav/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const VERSION = "1.0.0"

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	database.InitDB()

	gin.SetMode(viper.GetString("server.mode"))

	r := gin.Default()

	// Modify CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = viper.GetStringSlice("cors.allowOrigins") // Read from config file
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	setupRoutes(r)

	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)

		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.GET("/user", controllers.GetUserInfo)
			auth.PUT("/user/settings", controllers.UpdateUserSettings)

			auth.POST("/bookmarks", controllers.CreateBookmark)
			auth.GET("/bookmarks", controllers.GetBookmarks)
			auth.PUT("/bookmarks/:id", controllers.UpdateBookmark)
			auth.DELETE("/bookmarks/:id", controllers.DeleteBookmark)

			// Add categories routes
			auth.GET("/categories", controllers.GetCategories)
			auth.POST("/categories", controllers.CreateCategory)
			auth.PUT("/categories/:id", controllers.UpdateCategory)
			auth.DELETE("/categories/:id", controllers.DeleteCategory)

			auth.GET("/settings", controllers.GetSettings)
			auth.PUT("/settings", controllers.UpdateSettings)
		}

		// Public bookmark route
		api.GET("/share/:publicId", controllers.GetPublicBookmarks)
	}
}
