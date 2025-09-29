package restaurant

import (
	db "Food_Delivery_Management/DB"
	repository "Food_Delivery_Management/Repository"
	schema "Food_Delivery_Management/Schema"
	"Food_Delivery_Management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurants(ctx *gin.Context) {
	getUserId, _ := ctx.Get("user_id")
	getRole, _ := ctx.Get("role")

	if getRole != "Owner" {
		utils.RespondIfError(ctx, nil, http.StatusUnauthorized, "Owner role only allowed")
		return
	}

	if getUserId == nil {
		utils.RespondIfError(ctx, nil, http.StatusUnauthorized, "User id not found")
		return
	}
	option := &repository.QueryOptions{
		Limit:   10,
		Offset:  0,
		OrderBy: "id",
	}

	getRestaurantDB, getRestaurantDBError := repository.GetAllRecords[schema.Restaurants](db.DB, option)

	if getRestaurantDBError != nil {
		utils.RespondIfError(ctx, nil, http.StatusInternalServerError, "Failed to get restaurants")
		return
	}
	utils.HandleSuccess(ctx, http.StatusOK, getRestaurantDB)
}
