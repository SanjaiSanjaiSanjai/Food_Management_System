package register

import (
	db "Food_Delivery_Management/DB"
	"Food_Delivery_Management/DTO"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/crypto"
	"Food_Delivery_Management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterController(ctx *gin.Context) {
	var req DTO.RegisterDTO

	err := ctx.ShouldBind(&req)

	// error handling
	utils.RespondIfError(ctx, err, http.StatusBadRequest, "Invalid request body format from RegisterController function ShouldBind")
	customlogger.Log.Info("[RegisterController]: request body format is success")

	password := req.Password

	// hash password custom function
	hashPassword, hashErr := crypto.BcryptHash(password)

	// error handling
	utils.RespondIfError(ctx, hashErr, http.StatusInternalServerError, "BcryptHash is issue from RegisterController function BcryptHash")
	customlogger.Log.Info("[RegisterController]: BcryptHash is success")

	req.Password = string(hashPassword)

	user := schema.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	db_newUser_result, db_newUser_error := repository.CreateDB(db.DB, &user)

	// error handling
	utils.RespondIfError(ctx, db_newUser_error, http.StatusInternalServerError, "newuser create db error from RegisterController function CreateDB")
	customlogger.Log.Info("[RegisterController]: newuser create db is success")

	refreshToken, token_err := jwttoken.GenerateRefreshToken(db_newUser_result.ID, db_newUser_result.Email)

	// error handling
	utils.RespondIfError(ctx, token_err, http.StatusInternalServerError, "GenerateRefreshToken is error from RegisterController function GenerateRefreshToken")
	customlogger.Log.Info("[RegisterController]: GenerateRefreshToken is success")
	utils.HandleSuccess(ctx, http.StatusOK, gin.H{
		"message":       "register success",
		"refresh_token": refreshToken,
	})
}
