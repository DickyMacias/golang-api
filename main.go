package main

import (
	"html/template"
	"log"
	"movie-tracker/config"
	"movie-tracker/database"
	"movie-tracker/models"
	"movie-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	database.InitDatabase()

	// Auto migrate database tables
	db := database.GetDB()
	if err := db.AutoMigrate(&models.User{}, &models.FavoriteMovie{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create Gin router
	r := gin.Default()

	// Add custom template functions
	r.SetFuncMap(template.FuncMap{
		"seq": func(start, end int) []int {
			var result []int
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
			return result
		},
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Serve static files
	r.Static("/static", "./static")

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	log.Printf("🚀 Movie Tracker server starting on port %s", cfg.Port)
	log.Printf("🎬 Visit http://localhost:%s to access the application", cfg.Port)
	
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}