package restaurant

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

func CreateRestaurant(ctx *gin.Context) {

	getUserId, _ := ctx.Get("user_id")
	getRole, _ := ctx.Get("role")

	if getRole != "Owner" {
		utils.RespondIfError(ctx, "Owner role only allowed", http.StatusUnauthorized)
		return
	}

	if getUserId == nil {
		utils.RespondIfError(ctx, "User id not found", http.StatusUnauthorized)
		return
	}

	var role schema.Role

	findOneRoleDB, findOneRoleDBError := repository.FindOneWithConditions(db.DB, &role, []repository.QueryCondition{{Field: "id", Operator: "=", Value: getUserId}, {Field: "role", Operator: "=", Value: getRole}}, nil)

	if findOneRoleDBError != nil {
		customlogger.Log.Error("Failed to find role")
		utils.RespondIfError(ctx, "Failed to find role", http.StatusInternalServerError)
		return
	}
	customlogger.Log.Info("Role found")
	var req DTO.RestaurantDTO

	requestBodyDataError := ctx.ShouldBind(&req)

	fmt.Println("findOneRoleDB", findOneRoleDB)
	fmt.Println("role", role)
	fmt.Println("req", req)
	fmt.Println("requestBodyDataError", requestBodyDataError)
	// error handling
	utils.RespondIfError(ctx, requestBodyDataError, http.StatusBadRequest)

	createRestaurant := schema.Restaurants{
		Name:           req.Name,
		Description:    req.Description,
		Rating:         req.Rating,
		Cuisine_type:   req.Cuisine_type,
		Phone:          req.Phone,
		Email:          req.Email,
		License_number: req.License_number,
		Owner_id:       uint(getUserId.(uint)),
		Status:         true,
	}

	createRestaurantDB, createRestaurantDBError := repository.CreateDB(db.DB, &createRestaurant)
	if createRestaurantDBError != nil {
		customlogger.Log.Error("Failed to create restaurant")
		fmt.Println("Failed to create restaurant:", createRestaurantDBError)
		utils.RespondIfError(ctx, "Failed to create restaurant", http.StatusInternalServerError)
		return
	}
	customlogger.Log.Info("Restaurant created")
	utils.HandleSuccess(ctx, http.StatusOK, createRestaurantDB)
}
