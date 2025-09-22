package login

import (
	db "Food_Delivery_Management/DB"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	var user schema.User

	err := ctx.ShouldBind(&user)
	utils.IsNotNilError(err, "LoginController", "request body is error")

	// Method 1: Using the convenient FindByEmail function
	result, err := repository.FindByEmail(db.DB, &user, user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "User not found",
			"message": "Invalid email or password",
		})
		return
	}

	accessToken, err := jwttoken.GenerateAccessToken(result.ID, result.Email, "User")

	utils.IsNotNilError(err, "LoginController", "GenerateAccessToken is error")

	updateQueryErr := repository.UpdateDynamic(db.DB, &result, []repository.QueryCondition{{Field: "is_verified", Operator: "=", Value: false}}, map[string]interface{}{"IsVerified": true})

	utils.IsNotNilError(updateQueryErr, "LoginController", "UpdateWithConditions")
	// Method 2: Alternative using FindOneWithConditions for more complex queries
	// This demonstrates how you can use multiple conditions and options
	/*
		conditions := []repository.QueryCondition{
			{Field: "email", Operator: "=", Value: user.Email},
			{Field: "status", Operator: "=", Value: true}, // Only active users
			{Field: "is_verified", Operator: "=", Value: true}, // Only verified users
		}

		options := &repository.QueryOptions{
			Preload: []string{"Role", "User_Addresses"}, // Eager load relationships
			OrderBy: "created_at DESC",
		}

		result, err := repository.FindOneWithConditions(db.DB, &user, conditions, options)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "User not found",
				"message": "Invalid email or password",
			})
			return
		}
	*/

	// TODO: Add password verification here
	// Example: bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	// User found and authenticated, proceed with login logic
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":           result.ID,
			"username":     result.Username,
			"email":        result.Email,
			"IsVerified":   result.IsVerified,
			"access_Token": accessToken,
			// Don't return password in response
		},
	})
}
