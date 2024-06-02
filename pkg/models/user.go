package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"not null"`
	Password    string `gorm:"not null"`
}

type RegisterUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"creation_date"`
}
