package jwttoken

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	"Food_Delivery_Management/utils"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyJwtToken(tokenString string) (*jwt.Token, error) {
	godotenv.Load()

	// Get SECRET_KEY after loading .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	SECRET_KEY := []byte(secretKey)

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// Ensure the signing method is HMAC with HS256
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return SECRET_KEY, nil
	})

	// Handle parse errors
	if err != nil {
		utils.IsNotNilError(err, "VerifyJwtToken", "jwt Parse is issue")
		return nil, err
	}

	fmt.Println("token jwt.parse: ", token)
	utils.IsNotNilSuccess(token, "VerifyJwtToken", "jwt parse is token success")
	// Check if token is valid (avoid nil pointer panic)

	// fmt.Println("token jwt.valid: ", token.Valid)
	if token == nil || !token.Valid {
		customlogger.Log.Error("[VerifyJwtToken]: Invalid or expired token")
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}
