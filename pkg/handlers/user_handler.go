package handlers

import (
	"Dota2Api/initializers"
	"Dota2Api/pkg/models"
	"Dota2Api/pkg/wrappers"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	response := models.UserDto{
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

	response := make([]models.UserDto, len(users))
	for i, user := range users {
		response[i] = models.UserDto{
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
