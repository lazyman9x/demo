package auth

import (
	"github.com/lazyman9x/demo/v1/model"
	"github.com/gin-gonic/gin"
)

func RegisterAuthentication(ctx *gin.Context, user *model.User, userId uint, tokenId string) {
	ctx.Set("user", user)
	ctx.Set("tokenId", tokenId)
	ctx.Set("userId", userId)
}

func GetUserId(ctx *gin.Context) uint {
	user := ctx.MustGet("user").(*model.User)
	if user == nil {
		userId := ctx.MustGet("userId").(uint)
		if userId == 0 {
			panic("This user not exists")
		}
		return userId
	}
	return user.ID
}

func GetTokenId(ctx *gin.Context) string {
	return ctx.MustGet("tokenId").(string)
}
