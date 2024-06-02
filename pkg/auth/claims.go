package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email              string `json:"email"`
	jwt.StandardClaims `json:"standardClaims"`
}

type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
