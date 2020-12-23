package api

import (
	"github.com/gin-gonic/gin"
)

func successOrAbort(ctx *gin.Context, statusCode int, err error) bool {
	if err != nil {
		ctx.AbortWithError(statusCode, err)
	}
	return true
}
