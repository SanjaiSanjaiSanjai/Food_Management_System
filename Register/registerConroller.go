package register

import (
	db "Food_Delivery_Management/DB"
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
	customlogger.Log.Info("[RegisterController]: function call is received")
	// assign User Schema
	var user schema.User

	//get data req body
	err := ctx.ShouldBind(&user)

	// error throw if ShouldBind is return error
	utils.IsNotNilError(err, "RegisterController", "request body error")

	// success throw if ShouldBind is return data
	utils.IsNillSuccess(err, "RegisterController", "request body is success")

	// get password only in user
	password := user.Password

	// convert string password to hashpassword
	hashPassword, hashErr := crypto.BcryptHash(password)

	// error throw if BcryptHash is return error
	utils.IsNotNilError(hashErr, "RegisterController", "hashPassword is issue")

	// success throw if BcryptHash is return data
	utils.IsNotNilSuccess(hashPassword, "RegisterController", "hashPassword is genrated is succes")

	// convert byte to string password
	user.Password = string(hashPassword)

	//create new user DB function
	db_newUser_result, db_newUser_error := repository.CreateDB(db.DB, &user)

	//error throw if CreateDB is return error
	utils.IsNotNilError(db_newUser_error, "RegisterController", "newuser create db error")

	//success throw if CreateDB is return data
	utils.IsNotNilSuccess(db_newUser_result, "RegisterController", "newuser create db is succes")

	// Token Generated
	refreshToken, token_err := jwttoken.GenerateRefreshToken(db_newUser_result.ID, db_newUser_result.Email)

	//error throw if GenerateRefreshToken is return error
	utils.IsNotNilError(token_err, "RegisterController", "GenerateRefreshToken is error")

	// success throw if GenerateRefreshToken is return data
	utils.IsNotNilSuccess(refreshToken, "RegisterController", "refreshToken created  is succes")

	ctx.JSON(http.StatusOK, gin.H{"message": "register success", "refresh_token": refreshToken})
}
