package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var key = []byte(os.Getenv("TOKEN_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerPrefix = "Bearer "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization token not provided"})
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		// Check if token is valid
		validation, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if validation.Valid == false && err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
