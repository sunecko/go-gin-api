package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId     uint   `json:"userId"`
	Email      string `json:"email"`
	jwt.Claims `json:"claims"`
}

type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
