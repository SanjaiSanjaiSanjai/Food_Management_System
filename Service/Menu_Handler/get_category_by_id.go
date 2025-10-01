package menuhandler

import (
	// schema "Food_Delivery_Management/Schema"
	// "Food_Delivery_Management/utils"
	db "Food_Delivery_Management/DB"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenuCategoryById(ctx *gin.Context) {
	getRole, _ := ctx.Get("role")
	if getRole != "Owner" && getRole != "User" {
		utils.RespondIfError(ctx, "Unauthorized", http.StatusUnauthorized)
		return
	}
	category_id := ctx.Param("category_id")
	if category_id == "" {
		customlogger.Log.Error("Category ID is required")
		utils.RespondIfError(ctx, "Category ID is required", http.StatusBadRequest)
		return
	}
	customlogger.Log.Info("Category ID is received")
	var categorySchema schema.MenuCategory
	var category_condition = []repository.QueryCondition{{Field: "id", Operator: "=", Value: category_id}}

	is_available_categoryBy_id, not_available_categoryBy_id := repository.FindOneWithConditions(db.DB, &categorySchema, category_condition, nil)
	if not_available_categoryBy_id != nil {
		customlogger.Log.Error("Category not found")
		utils.RespondIfError(ctx, not_available_categoryBy_id, http.StatusNotFound)
		return
	}
	customlogger.Log.Info("Category found")
	fmt.Println("is_available_categoryBy_id", is_available_categoryBy_id)

	var restaurantCategory []schema.RestaurantCategory
	var restaurant_category_condition = []repository.QueryCondition{{Field: " menu_category_id ", Operator: "=", Value: category_id}}
	getAll_restaurant_categories, getAll_restaurant_categoriesErr := repository.FindManyWithConditions(db.DB, &restaurantCategory, restaurant_category_condition, nil)

	if getAll_restaurant_categoriesErr != nil {
		customlogger.Log.Error("Failed to get restaurant categories")
		utils.RespondIfError(ctx, getAll_restaurant_categoriesErr, http.StatusInternalServerError)
		return
	}
	customlogger.Log.Info("Restaurant categories found")
	utils.HandleSuccess(ctx, http.StatusOK, getAll_restaurant_categories)
}
