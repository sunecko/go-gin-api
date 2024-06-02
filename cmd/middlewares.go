package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
