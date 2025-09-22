package jwttoken

import (
	"Food_Delivery_Management/utils"
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

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateRefreshToken(user_id uint, email string) (string, error) {
	godotenv.Load()
	claims := &Claims{
		UserID: user_id,
		Email:  email,
		Role:   "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Food_Delivery_Management",
			Subject:   email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SECRET_KEY)
	utils.IsNotNilError(err, "GenerateRefreshToken", "tokenString is issue")
	return tokenString, nil
}
