package handlers

import (
	"Dota2Api/initializers"
	"Dota2Api/pkg/auth"
	"Dota2Api/pkg/models"
	"Dota2Api/pkg/wrappers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	req := models.RegisterUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error binding request",
			Detail: "Invalid request body"})
		return
	}

	_, err = auth.FindUserByEmail(req.Email)
	if err == nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error creating user",
			Detail: "User with this email already exists"})
		return
	}

	password := auth.EncryptPassword(req.Password)

	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    password,
	}

	dbCall := initializers.DB.Create(&user)
	if dbCall.Error != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error creating user",
			Detail: fmt.Sprintf("Error: %v", dbCall.Error),
		})
		return
	}

	c.JSON(201, wrappers.CreatedResponse{
		Id:      user.ID,
		Message: "User created successfully",
	})
}

func Login(c *gin.Context) {
	req := models.LoginUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error binding request",
			Detail: "Invalid request body"})
		return
	}

	passwordValidation, _ := auth.VerifyPassword(req.Email, req.Password)
	if passwordValidation == false {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Credentials errors",
			Detail: "Incorrect email or password"})
	}

	response, err := auth.CreateToken(req.Email)
	if err != nil {
		c.JSON(400, wrappers.ResponseError{
			Title:  "Error creating token",
			Detail: fmt.Sprintf("Error: %v", err)})
		return
	}

	c.JSON(200, response)
}
