package menuhandler

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

/*
create menu
@route POST /v1/api/menu/create/:restaurant_id
@access Owner only
@body {"name":"string","description":"string","price":"number","is_vegetarian":"boolean"}
@getRole --> get role from context
@restaurantID --> get restaurant_id from params
@findRestaurant --> find restaurant by restaurant_id
@restaurant_condition --> create restaurant condition
@restaurant_table --> restaurant schema type
@menuDTO --> menu dto type
@newMenu --> create new menu
@categoryCondition --> create category condition
@findCategory --> find category by category_name
@menuCategory --> menu category schema type
@requestBodyErr --> check request body error
@utils.HandleSuccess --> handle success response
@utils.RespondIfError --> handle error response

*/

func CreateMenu(ctx *gin.Context) {
	getRole, ok := ctx.Get("role")
	var restaurant_table schema.Restaurants
	var menuDTO DTO.MenuDTO
	var menuCategory schema.MenuCategory

	fmt.Println("getRole", getRole)
	if !ok || getRole != "Owner" {
		customlogger.Log.Error("Owner role only allowed")
		utils.RespondIfError(ctx, "Owner role only allowed", http.StatusUnauthorized)
		return
	}
	customlogger.Log.Info("Owner role allowed")
	restaurantID := ctx.Param("restaurant_id")
	fmt.Println("restaurantID", restaurantID)

	restaurant_condition := []repository.QueryCondition{
		{Field: "id", Operator: "=", Value: restaurantID},
	}
	findRestaurant, findRestaurantErr := repository.FindOneWithConditions(db.DB, &restaurant_table, restaurant_condition, nil)
	if findRestaurantErr != nil {
		utils.RespondIfError(ctx, findRestaurantErr, http.StatusNotFound)
	}

	requestBodyErr := ctx.ShouldBind(&menuDTO)
	if requestBodyErr != nil {
		utils.RespondIfError(ctx, requestBodyErr, http.StatusBadRequest)
	}

	categoryCondition := []repository.QueryCondition{
		{Field: "name", Operator: "=", Value: menuDTO.CategoryName},
	}
	findCategory, findCategoryErr := repository.FindOneWithConditions(db.DB, &menuCategory, categoryCondition, nil)
	if findCategoryErr != nil {
		utils.RespondIfError(ctx, findCategoryErr, http.StatusNotFound)
	}

	// create new menu
	newMenu := schema.Menu{
		Name:         menuDTO.Name,
		Description:  menuDTO.Description,
		Price:        menuDTO.Price,
		IsVegetarian: menuDTO.IsVegetarian,
		CategoryID:   findCategory.ID,
		RestaurantID: findRestaurant.Id,
	}
	fmt.Println("newMenu", newMenu)
	createNewMenu, createNewMenuErr := repository.CreateDB(db.DB, &newMenu)
	if createNewMenuErr != nil {
		utils.RespondIfError(ctx, createNewMenuErr, http.StatusInternalServerError)
	}
	utils.HandleSuccess(ctx, http.StatusCreated, createNewMenu)
}
