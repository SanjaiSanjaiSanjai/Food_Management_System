package useraddress

import (
	db "Food_Delivery_Management/DB"
	"Food_Delivery_Management/DTO"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAddressController(ctx *gin.Context) {
	var request_user_address DTO.UserAddressDTO

	getUserId, _ := ctx.Get("user_id")
	getRole, _ := ctx.Get("role")

	if getRole != "User" {
		customlogger.Log.Error("User role is not valid")
		utils.RespondIfError(ctx, "User role is not valid", http.StatusUnauthorized)
		return
	}

	requestBodyError := ctx.ShouldBindJSON(&request_user_address)

	if requestBodyError != nil {
		customlogger.Log.Error("Error in binding request body")
		fmt.Println("Error in binding request body:", requestBodyError)
		utils.RespondIfError(ctx, requestBodyError, http.StatusBadRequest)
		return
	}

	customlogger.Log.Info("Request body binding successfully")
	user_address_DB := schema.User_Addresses{
		UserID:     getUserId.(uint),
		Address:    request_user_address.Address,
		State:      request_user_address.State,
		Country:    request_user_address.Country,
		Postalcode: request_user_address.Postalcode,
		Landmark:   request_user_address.Landmark,
		Status:     request_user_address.Status,
	}

	create_user_address_DB, create_user_address_DBError := repository.CreateDB(db.DB, &user_address_DB)
	if create_user_address_DBError != nil {
		customlogger.Log.Error("Error in creating user address")
		fmt.Println("Error in creating user address:", create_user_address_DBError)
		utils.RespondIfError(ctx, "Error in creating user address", http.StatusInternalServerError)
		return
	}
	utils.HandleSuccess(ctx, http.StatusOK, create_user_address_DB)
}
