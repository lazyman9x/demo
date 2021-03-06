package router

import (
	"github.com/lazyman9x/demo/v1/api"
	"github.com/lazyman9x/demo/v1/config"
	"github.com/lazyman9x/demo/v1/database"
	"github.com/lazyman9x/demo/v1/model"
	"github.com/gin-gonic/gin"
)

func Create(db *database.GormDatabase, vInfo *model.VersionInfo, conf *config.Configuration) (*gin.Engine, func()) {
	g := gin.New()
	g.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		for header, value := range conf.Server.ResponseHeaders {
			ctx.Header(header, value)
		}
	})
	clientHandler := api.ClientAPI{
		DB:       db,
		ImageDir: conf.UploadedImagesDir,
	}

	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	clientAuth := g.Group("")
	{
		client := clientAuth.Group("/client")
		{
			client.POST("", clientHandler.CreateClient)
		}
	}

	userHandler := api.UserAPI{DB: db, PasswordStrength: conf.PassStrength}
	userAuth := g.Group("/user")
	{
		userAuth.POST("", userHandler.CreateUser)
	}
	return g, nil
}
