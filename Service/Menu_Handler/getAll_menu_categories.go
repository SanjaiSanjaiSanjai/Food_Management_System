package menuhandler

import (
	db "Food_Delivery_Management/DB"
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMenuCategories(ctx *gin.Context) {
	getRole, _ := ctx.Get("role")
	if getRole != "Owner" && getRole != "User" {
		customlogger.Log.Error("Unauthorized from GetAllMenuCategories")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	customlogger.Log.Info("GetAllMenuCategories is called")
	var menuCategories []schema.MenuCategory

	queryCondition := []repository.QueryCondition{
		{Field: "is_active", Operator: "=", Value: true},
		{Field: "status", Operator: "=", Value: true},
	}
	getMenuCategories, err := repository.FindManyWithConditions(db.DB, &menuCategories, queryCondition, nil)
	if err != nil {
		customlogger.Log.Error("Failed to get menu categories from GetAllMenuCategories in DB")
		utils.RespondIfError(ctx, "failed to get menu categories", http.StatusInternalServerError)
		return
	}
	customlogger.Log.Info("GetAllMenuCategories is success")
	utils.HandleSuccess(ctx, http.StatusOK, getMenuCategories)
}
