package restaurant

import (
	db "Food_Delivery_Management/DB"
	"Food_Delivery_Management/DTO"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
CreateRestaurantAddress
only owner can create restaurant address
if owner try to create restaurant address then it will create after that it will return restaurant details
*/

func CreateRestaurantAddress(ctx *gin.Context) {
	getUserId, _ := ctx.Get("user_id")
	getRole, _ := ctx.Get("role")
	var req DTO.RestaurantAddressDTO
	var restaurant schema.Restaurants

	if getRole != "Owner" {
		utils.RespondIfError(ctx, nil, http.StatusUnauthorized, "Owner role only allowed")
		return
	}
	customlogger.Log.Info("Owner role allowed")
	if getUserId == nil {
		utils.RespondIfError(ctx, nil, http.StatusUnauthorized, "User id not found")
		return
	}
	customlogger.Log.Info("User id found")

	condition := []repository.QueryCondition{
		{Field: "owner_id", Operator: "=", Value: getUserId},
	}
	options := &repository.QueryOptions{
		OrderBy: "created_at DESC", // or any other ordering you prefer
	}
	getRestaurantDB, getRestaurantDBError := repository.FindOneWithConditions(db.DB, &restaurant, condition, options)

	if getRestaurantDBError != nil {
		customlogger.Log.Error("Failed to find restaurant")
		utils.RespondIfError(ctx, nil, http.StatusInternalServerError, "Failed to find restaurant")
		return
	}

	if getRestaurantDB == nil {
		customlogger.Log.Error("Restaurant not found")
		utils.RespondIfError(ctx, nil, http.StatusNotFound, "Restaurant not found")
		return
	}

	customlogger.Log.Info("Restaurant found")

	requestBodyErr := ctx.ShouldBindJSON(&req)
	if requestBodyErr != nil {
		utils.RespondIfError(ctx, nil, http.StatusBadRequest, "Bad request")
		return
	}

	createRestaurantAddress := schema.RestaurantAddress{
		RestaurantID: getRestaurantDB.Id,
		Address:      req.Address,
		City:         req.City,
		State:        req.State,
		PostalCode:   req.Postalcode,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
	}
	createdAddress, createRestaurantAddressError := repository.CreateDB(db.DB, &createRestaurantAddress)

	if createRestaurantAddressError != nil {
		customlogger.Log.Error("Failed to create restaurant address")
		utils.RespondIfError(ctx, nil, http.StatusInternalServerError, "Failed to create restaurant address")
		return
	}
	customlogger.Log.Info("Restaurant address created successfully")
	utils.HandleSuccess(ctx, http.StatusCreated, createdAddress)
}
