package jwttoken

import (
	"Food_Delivery_Management/utils"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID uint   `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// SECRET_KEY will be loaded dynamically to ensure .env is loaded first

func GenerateRefreshToken(user_id uint, email string) (string, error) {
	godotenv.Load()
	
	// Get SECRET_KEY after loading .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	SECRET_KEY := []byte(secretKey)
	utils.IsNillError(SECRET_KEY, "GenerateRefreshToken", "SECRET_KEY is nil")

	claims := &Claims{
		UserID: user_id,
		Email:  email,
		Role:   "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Food_Delivery_Management",
			Subject:   "email",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	utils.IsNillError(token, "GenerateRefreshToken", "token is NewWithClaims issue")

	tokenString, err := token.SignedString(SECRET_KEY)
	utils.IsNotNilError(err, "GenerateRefreshToken", "tokenString is issue")
	return tokenString, nil
}
