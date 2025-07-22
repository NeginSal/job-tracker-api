package main

import (
	"log"

	"github.com/NeginSal/job-tracker-api/internal/middleware"
	"github.com/NeginSal/job-tracker-api/internal/user"
	"github.com/NeginSal/job-tracker-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load .env file")
	}

	database := db.NewPostgresConnection()
	database.AutoMigrate(&user.User{})

	userRepo := user.NewRepository(database)
	userService := user.NewService(userRepo)
	userHandler := *user.NewHandler(userService)

	r := gin.Default()

	// public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/protected", func(c *gin.Context) {
		userID := c.MustGet("userID").(string)
		c.JSON(200, gin.H{"message": "Access granted", "user_id": userID})
	})

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}