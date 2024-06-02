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
	router := gin.New()

	//Use the logger and recovery middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Migrate the User model
	migrationErr := initializers.DB.AutoMigrate(&models.User{})
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	needAuth := router.Group("/")

	//Auth routes
	router.POST("/auth/register", handlers.RegisterUser)
	router.POST("/auth/login", handlers.Login)

	needAuth.Use(AuthMiddleware())
	{
		router.DELETE("/user/:id", handlers.DeleteUser)
		router.GET("/user/:id", handlers.GetUser)
		router.GET("/user", AuthMiddleware(), handlers.GetAllUsers)
	}

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
