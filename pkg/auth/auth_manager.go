package auth

import (
	"Dota2Api/initializers"
	"Dota2Api/pkg/models"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey = os.Getenv("JWT_KEY")

func FindUserByEmail(email string) (models.User, error) {
	user := models.User{}
	result := initializers.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func FindUserById(id uint) (models.User, error) {
	user := models.User{}
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func EncryptPassword(password string) string {
	encryptedPassword := sha256.Sum256([]byte(password))
	encryptedPasswordStr := hex.EncodeToString(encryptedPassword[:])

	return encryptedPasswordStr
}

func VerifyPassword(email string, password string) (bool, error) {
	user, err := FindUserByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Password == EncryptPassword(password) {
		return true, nil
	}

	return false, nil
}

func CreateToken(email string) (TokenResponse, error) {
	expirationTime := time.Now().Add(1440 * time.Minute * 30)

	claims := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		Token:     tokenString,
		ExpiresAt: expirationTime.Unix(),
	}, nil
}
