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

	//User routes
	router.DELETE("/user/:id", handlers.DeleteUser)
	router.GET("/user/:id", handlers.GetUser)
	router.GET("/user", handlers.GetAllUsers)

	//Auth routes
	router.POST("/auth/register", handlers.RegisterUser)
	router.POST("/auth/login", handlers.Login)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
