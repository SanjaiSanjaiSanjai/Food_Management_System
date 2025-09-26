package restaurant

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(router *gin.RouterGroup) {
	router.POST(baseurl.RESTAURANT_URL["restaurant"], jwttoken.ExtractTokenFromRequest, CreateRestaurant)
}
