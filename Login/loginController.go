package login

import (
	db "Food_Delivery_Management/DB"
	"Food_Delivery_Management/DTO"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"
	"Food_Delivery_Management/crypto"
	"Food_Delivery_Management/utils"

	// customlogger "Food_Delivery_Management/HandleCustomLogger"

	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"

	// "Food_Delivery_Management/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	var req DTO.LoginDTO
	var user schema.User
	var role schema.Role

	getTokenUserId, ok := ctx.Get("user_id")

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format in LoginController function"})
		ctx.Abort()
		return
	}

	requestBodyDataError := ctx.ShouldBind(&req)

	// error handling
	utils.RespondIfError(ctx, requestBodyDataError, http.StatusBadRequest, "Invalid request body format in LoginController function")

	getUserDataById, userFindError := repository.FindByID(db.DB, &user, getTokenUserId.(uint))

	// error handling
	utils.RespondIfError(ctx, userFindError, http.StatusUnauthorized, "User not found")
	customlogger.Log.Info("[LoginController]: User found")
	fmt.Printf("getUserDataById: %v", getUserDataById)

	if req.Email != getUserDataById.Email {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		ctx.Abort()
		return
	}

	isValid, compareError := crypto.BcryptCompare([]byte(getUserDataById.Password), req.Password)

	// error handling
	utils.RespondIfError(ctx, compareError, http.StatusUnauthorized, "Invalid password")

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		ctx.Abort()
		return
	}
	customlogger.Log.Info("[LoginController]: Password matched")

	accessToken, generateAccessTokenError := jwttoken.GenerateAccessToken(getUserDataById.ID, getUserDataById.Email, req.Role)

	// error handling
	utils.RespondIfError(ctx, generateAccessTokenError, http.StatusUnauthorized, "Invalid authorization header format")

	updateQueryErr := repository.UpdateWithConditions(db.DB, &getUserDataById, []repository.QueryCondition{{Field: "is_verified", Operator: "=", Value: false}}, map[string]interface{}{"IsVerified": true})

	// error handling
	utils.RespondIfError(ctx, updateQueryErr, http.StatusUnauthorized, "Invalid authorization header format")
	customlogger.Log.Info("[LoginController]: UpdateWithConditions is success")
	role = schema.Role{UserID: getUserDataById.ID, Role: req.Role, Status: true}
	createRole, createRoleError := repository.CreateDB(db.DB, &role)

	// error handling
	utils.RespondIfError(ctx, createRoleError, http.StatusUnauthorized, "Invalid authorization header format")
	customlogger.Log.Info("[LoginController]: CreateDB is success")

	utils.HandleSuccess(ctx, http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":           getUserDataById.ID,
			"username":     getUserDataById.Username,
			"email":        getUserDataById.Email,
			"IsVerified":   getUserDataById.IsVerified,
			"role":         createRole.Role,
			"access_Token": accessToken,
		},
	})
}
