package menuhandler

import (
	// schema "Food_Delivery_Management/Schema"
	// "Food_Delivery_Management/utils"
	db "Food_Delivery_Management/DB"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
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
		utils.RespondIfError(ctx, "Category not found", http.StatusNotFound)
		return
	}
	customlogger.Log.Info("Category found")
	utils.HandleSuccess(ctx, http.StatusOK, is_available_categoryBy_id)
}
