package handlers

import (
	"Dota2Api/initializers"
	"Dota2Api/pkg/models"
	"Dota2Api/pkg/wrappers"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"creation_date"`
}

func CreateUser(c *gin.Context) {
	req := CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error binding request",
			Detail: "Invalid request body"})
		return
	}

	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, user.ID)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error binding request",
			Detail: "Invalid request body"})
		return
	}

	user := models.User{}
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, wrappers.ResponseError{
			Title:  "Error fetching user",
			Detail: "User not found"})
	}

	user.Name = req.Name
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber
	user.UpdatedAt = time.Now()

	result = initializers.DB.Save(&user)

	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	response := UserDto{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.String(),
	}

	c.JSON(200, response)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(404, wrappers.ResponseError{
			Title:  "Error fetching users",
			Detail: "No users found"})
		return
	}

	response := make([]UserDto, len(users))
	for i, user := range users {
		response[i] = UserDto{
			Id:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			CreatedAt:   user.CreatedAt.String(),
		}
	}

	c.JSON(200, response)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}

	result := initializers.DB.Delete(&user, id)

	if result.Error != nil {
		c.JSON(404, wrappers.ResponseError{
			Title:  "Error deleting user",
			Detail: "User not found"})
		return
	}

	c.JSON(204, "User deleted successfully")
}
