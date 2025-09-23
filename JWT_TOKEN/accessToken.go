package jwttoken

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	// "Food_Delivery_Management/utils"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateAccessToken(user_id uint, email string, role string) (string, error) {
	godotenv.Load()

	// Get SECRET_KEY after loading .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	SECRET_KEY := []byte(secretKey)

	claims := &Claims{
		UserID: user_id,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Food_Delivery_Management",
			Subject:   "email",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		customlogger.Log.Error("[GenerateAccessToken]: SignedString is error")
		return "", err
	}

	return tokenString, err
}
