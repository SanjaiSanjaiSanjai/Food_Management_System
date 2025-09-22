package jwttoken

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	"Food_Delivery_Management/utils"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwtToken(tokenString string) (*jwt.Token, error) {
	var SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	if SECRET_KEY == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	utils.IsNotNilError(err, "VerifyJwtToken", "jwt Parse is issue")

	if !token.Valid {
		customlogger.Log.Error("[VerifyJwtToken]:Invalid or expired token")
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}
