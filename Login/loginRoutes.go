package login

import (
	baseurl "Food_Delivery_Management/BaseUrl"
	jwttoken "Food_Delivery_Management/JWT_TOKEN"

	"github.com/gin-gonic/gin"
)

func LoginRouter(rg *gin.RouterGroup) {
	rg.POST(baseurl.LOGIN_URL["login"], jwttoken.ExtractTokenFromRequest, LoginController)
}
