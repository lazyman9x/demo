package api

import (
	"github.com/demo/server/v2/auth"
	"github.com/demo/server/v2/model"
	"github.com/gin-gonic/gin"
)

// The ClientDatabase interface for encapsulating database access.
type ClientDatabase interface {
	CreateClient(client *model.Client) error
	GetClientByToken(token string) (*model.Client, error)
	GetClientByID(id uint) (*model.Client, error)
	GetClientsByUser(userID uint) ([]*model.Client, error)
	DeleteClientByID(id uint) error
	UpdateClient(client *model.Client) error
}

// The ClientAPI provides handlers for managing clients and applications.
type ClientAPI struct {
	DB       ClientDatabase
	ImageDir string
}

func (a *ClientAPI) CreateClient(ctx *gin.Context) {
	client := model.Client{}
	if err := ctx.Bind(&client); err == nil {
		client.Token = auth.GenerateNotExistingToken(generateClientToken, a.isClientExists)
		client.UserID = auth.GetUserId(ctx)
		if success := successOrAbort(ctx, 500, a.DB.CreateClient(&client)); !success {
			return
		}
		ctx.JSON(200, client)
	}
}

func (a *ClientAPI) isClientExists(token string) bool {
	client, _ := a.DB.GetClientByToken(token)
	return client != nil
}
