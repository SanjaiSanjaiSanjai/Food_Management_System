package menuhandler

import (
	db "Food_Delivery_Management/DB"
	"Food_Delivery_Management/DTO"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Create Menu Category
@route POST /menu/category/create/:restaurant_id --> get restaurant_id from url
@request body {"name":"","description":"","is_active":true}
@menuCategory schema.MenuCategory <--- database table
@menuCategoryDTO DTO.MenuCategoryDTO <--- request body
@restaurantID uint <--- restaurant_id from url
@getRole string <--- role from context
@newCategory *schema.MenuCategory <--- created menu category
@createCategoryErr error <--- error from CreateDB
*/
func CreateMenuCategory(ctx *gin.Context) {
	restaurantIDStr := ctx.Param("restaurant_id")
	restaurantID, restaurantIDErr := strconv.ParseUint(restaurantIDStr, 10, 32)

	if restaurantIDErr != nil {
		utils.RespondIfError(ctx, "Bad request", http.StatusBadRequest)
	}
	var menuCategoryDTO DTO.MenuCategoryDTO

	getRole, _ := ctx.Get("role")
	if getRole != "Owner" {
		customlogger.Log.Error("Owner role only allowed")
		utils.RespondIfError(ctx, "Owner role only allowed", http.StatusUnauthorized)
		return
	}

	requestBodyErr := ctx.ShouldBind(&menuCategoryDTO)
	if requestBodyErr != nil {
		utils.RespondIfError(ctx, "Bad request", http.StatusBadRequest)
	}

	menuCategory := schema.MenuCategory{
		Name:         menuCategoryDTO.Name,
		Description:  menuCategoryDTO.Description,
		IsActive:     menuCategoryDTO.IsActive,
		RestaurantID: uint(restaurantID),
	}

	// Create menu category in database
	newCategory, createCategoryErr := repository.CreateDB(db.DB, &menuCategory)
	if createCategoryErr != nil {
		customlogger.Log.Error("Failed to create menu category from CreateDB")
		utils.RespondIfError(ctx, "Internal server error", http.StatusInternalServerError)
	}
	customlogger.Log.Info("Menu category created successfully")
	utils.HandleSuccess(ctx, http.StatusCreated, newCategory)
}
