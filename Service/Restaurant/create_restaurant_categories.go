package restaurant

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

func CreateRestaurantCategories(ctx *gin.Context) {
	getRole, _ := ctx.Get("role")
	if getRole != "Owner" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Owner role only allowed"})
		return
	}
	restaurantIdStr := ctx.Param("restaurant_id")
	restaurantId, err := strconv.ParseUint(restaurantIdStr, 10, 64)
	if err != nil {
		utils.RespondIfError(ctx, err, http.StatusBadRequest)
		return
	}

	var req DTO.RestaurantCategoryDTO
	requestBodyErr := ctx.ShouldBind(&req)
	if requestBodyErr != nil {
		utils.RespondIfError(ctx, requestBodyErr, http.StatusBadRequest)
		return
	}

	// get category id by category name
	var categorySchema schema.MenuCategory
	categoryCondition := []repository.QueryCondition{
		{Field: "name", Operator: "=", Value: req.CategoryName},
	}
	getCategoryIdByCategoryName, getCategoryIdByCategoryNameErr := repository.FindOneWithConditions(db.DB, &categorySchema, categoryCondition, nil)
	if getCategoryIdByCategoryNameErr != nil {
		utils.RespondIfError(ctx, getCategoryIdByCategoryNameErr, http.StatusInternalServerError)
		return
	}

	createNewRestaurantCategory := schema.RestaurantCategory{
		RestaurantID: uint(restaurantId),
		CategoryID:   getCategoryIdByCategoryName.ID,
	}

	newRestaurantCategory, newRestaurantCategoryErr := repository.CreateDB(db.DB, &createNewRestaurantCategory)
	if newRestaurantCategoryErr != nil {
		customlogger.Log.Error("Failed to create restaurant category")
		utils.RespondIfError(ctx, newRestaurantCategoryErr, http.StatusInternalServerError)
		return
	}
	customlogger.Log.Info("Restaurant category created successfully")
	utils.HandleSuccess(ctx, http.StatusCreated, newRestaurantCategory)
}
