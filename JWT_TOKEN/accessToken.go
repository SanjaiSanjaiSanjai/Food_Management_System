package jwttoken

import (
	"Food_Delivery_Management/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user_id uint, email string, role string) (string, error) {
	claims := &Claims{
		UserID: user_id,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Food_Delivery_Management",
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	utils.IsNotNilError(err, "GenerateAccessToken", "tokenString")

	return tokenString, nil
}
