package useraddress

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"

	"github.com/gin-gonic/gin"
)

func UserAddressRoutes(router *gin.RouterGroup) {
	router.POST(baseurl.USER_ADDRESS_URL["user_address"], jwttoken.ExtractTokenFromRequest, UserAddressController)
}
