package menuhandler

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(rg *gin.RouterGroup) {
	// create menu
	rg.POST(baseurl.MENU_URL["menu"], jwttoken.ExtractTokenFromRequest, CreateMenu)
	// create menu category
	rg.POST(baseurl.MENU_URL["menu_category"], jwttoken.ExtractTokenFromRequest, CreateMenuCategory)
	// get all menu categories
	rg.GET(baseurl.MENU_URL["getMenuCategories"], jwttoken.ExtractTokenFromRequest, GetAllMenuCategories)
}
