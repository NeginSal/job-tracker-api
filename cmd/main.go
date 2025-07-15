package main

import (
	"log"

	"github.com/NeginSal/job-tracker-api/internal/user"
	"github.com/NeginSal/job-tracker-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"job-tracker-api/internal/user"
	"job-tracker-api/pkg/db"
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
	r.POST("/register", userHandler.Register)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
