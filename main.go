package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DB_URL")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := db.AutoMigrate(&University{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found — using environment variables only")
	}

	initDB()
	RunScraper(db)

	router := gin.Default()

	router.GET("/universities", func(c *gin.Context) {
		category := c.Query("category")
		var unis []University
		query := db

		if category != "" {
			query = query.Where("category = ?", category)
		}

		if err := query.Find(&unis).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, unis)
	})

	router.GET("/universities/:id", func(c *gin.Context) {
		id := c.Param("id")
		var uni University
		result := db.First(&uni, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "University not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, uni)
	})

	router.POST("/scrape", func(c *gin.Context) {
		go RunScraper(db)
		c.JSON(http.StatusOK, gin.H{"message": "Scraper started in background"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
