package register

import (
	baseurl "Food_Delivery_Management/BaseUrl"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(rg *gin.RouterGroup) {
	rg.POST(baseurl.REGISTER_URL["register"], RegisterController)
}
