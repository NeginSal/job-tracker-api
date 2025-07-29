package main

import (
	"log"

	"github.com/NeginSal/job-tracker-api/internal/job"
	"github.com/NeginSal/job-tracker-api/internal/middleware"
	"github.com/NeginSal/job-tracker-api/internal/user"
	"github.com/NeginSal/job-tracker-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env file")
	}

	database := db.NewPostgresConnection()

	// Auto migrate both user and job tables
	database.AutoMigrate(&user.User{}, &job.Job{})

	// User setup
	userRepo := user.NewRepository(database)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// Job setup
	jobRepo := job.NewRepository(database)
	jobService := job.NewService(jobRepo)
	jobHandler := job.NewHandler(jobService)

	r := gin.Default()

	// Public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/protected", func(c *gin.Context) {
		userID := c.MustGet("userID").(string)
		c.JSON(200, gin.H{"message": "Access granted", "user_id": userID})
	})
	jobHandler.RegisterRoutes(auth)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}