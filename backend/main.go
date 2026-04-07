package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database connection
	// DB = ConnectDatabase()

	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
