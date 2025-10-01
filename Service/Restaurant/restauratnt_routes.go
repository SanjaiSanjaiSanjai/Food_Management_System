package restaurant

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(router *gin.RouterGroup) {
	router.POST(baseurl.RESTAURANT_URL["restaurant"], jwttoken.ExtractTokenFromRequest, CreateRestaurant)
	router.GET(baseurl.RESTAURANT_URL["getRestaurant"], jwttoken.ExtractTokenFromRequest, GetRestaurants)
	router.POST(baseurl.RESTAURANT_URL["restaurantAddress"], jwttoken.ExtractTokenFromRequest, CreateRestaurantAddress)
	router.POST(baseurl.RESTAURANT_CATEGORY_URL["restaurant_category"], jwttoken.ExtractTokenFromRequest, CreateRestaurantCategories)
}
