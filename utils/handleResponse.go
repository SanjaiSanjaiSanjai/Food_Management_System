package utils

import (
	"github.com/gin-gonic/gin"
)

// HandleSuccess writes a JSON success response with the given data.
// Usage:
//
//	utils.HandleSuccess(ctx, http.StatusOK, map[string]interface{}{"message": "success"})
func HandleSuccess(ctx *gin.Context, httpStatus int, response interface{}) {
	ctx.JSON(httpStatus, gin.H{"data": response})
}

// RespondIfError checks if err is not nil, and if so, writes a JSON error response
// and returns true indicating the response has been written and the caller should return.
// Usage:
//
//	if utils.RespondIfError(ctx, err, http.StatusInternalServerError, "BcryptHash is issue") { return }
func RespondIfError(ctx *gin.Context, err error, httpStatus int, message string) bool {
	if err != nil {
		ctx.JSON(httpStatus, gin.H{"error": message})
		return true
	}
	return false
}
