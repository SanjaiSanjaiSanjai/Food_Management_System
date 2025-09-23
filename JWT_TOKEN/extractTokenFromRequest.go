package jwttoken

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractTokenFromRequest(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	split := strings.Split(authorization, " ")

	if len(split) != 2 || split[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format in split"})
		ctx.Abort()
		return
	}

	token := split[1]

	claims, err := VerifyJwtToken(token)

	if err != nil {
		customlogger.Log.Error("[ExtractTokenFromRequest]: VerifyJwtToken is error")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format in VerifyJwtToken function"})
		ctx.Abort()
		return
	}

	tokenClaims, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		customlogger.Log.Error("[ExtractTokenFromRequest]: claims.Claims.(jwt.MapClaims) is error")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format in claims.Claims.(jwt.MapClaims) function"})
		ctx.Abort()
		return
	}
	fmt.Printf("tokenClaims type %T", tokenClaims)
	fmt.Println("tokenClaims ******: ", tokenClaims)

	// Our custom Claims uses json tag `json:"id"`, and when parsed into MapClaims
	// numeric values are float64. Convert to uint and store in context.
	if idf, ok := tokenClaims["id"].(float64); ok {
		fmt.Printf("idf type %T", idf)
		fmt.Println("idf: ", idf)
		ctx.Set("user_id", uint(idf))
	} else {
		customlogger.Log.Error("[ExtractTokenFromRequest]: tokenClaims['id'] not found or not a number")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: user id missing"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
