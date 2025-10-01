package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
}

func HandleSuccess(ctx *gin.Context, httpStatus int, data interface{}) {
	ctx.JSON(httpStatus, Response{Success: true, Data: data, Message: "success"})
}

func RespondIfError(ctx *gin.Context, msg interface{}, httpStatus int) {
	switch typeOfMsg := msg.(type) {
	case error:
		ctx.JSON(httpStatus, Response{Success: false, Error: typeOfMsg.Error(), Message: typeOfMsg.Error()})
	case string:
		ctx.JSON(httpStatus, Response{Success: false, Error: typeOfMsg, Message: typeOfMsg})
	default:
		ctx.JSON(httpStatus, Response{Success: false, Error: "", Message: "error"})
	}
}
