package jwttoken

import (
	"Food_Delivery_Management/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractTokenFromRequest(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	split := strings.Split(authorization, " ")

	if len(split) == 0 || split[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		ctx.Abort()
		return
	}

	token := split[1]

	claims, err := VerifyJwtToken(token)

	utils.IsNotNilError(err, "ExtractTokenFromRequest", "VerifyJwtToken is error")

	if claims != nil {
		ctx.Set("jwt_token", token)

		ctx.Next()
	}
	// ctx.Set("JWT_Token",token)
	// ctx.Next()
}
