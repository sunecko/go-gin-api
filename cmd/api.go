package main

import (
	"Dota2Api/initializers"
	"Dota2Api/pkg/handlers"
	"Dota2Api/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	//Initialize the router
	router := gin.Default()

	//Migrate the User model
	migrationErr := initializers.DB.AutoMigrate(&models.User{})
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	router.POST("/user", handlers.CreateUser)
	router.PUT("/user/:id", handlers.UpdateUser)
	router.GET("/user/:id", handlers.GetUser)
	router.GET("/user", handlers.GetAllUsers)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
