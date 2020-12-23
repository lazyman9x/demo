package api

import (
	"errors"

	"github.com/demo/server/v2/auth/password"
	"github.com/demo/server/v2/model"
	"github.com/gin-gonic/gin"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	DeleteUserByID(id uint) error
	UpdateUser(user *model.User) error
	CreateUser(user *model.User) error
	CountUser(condition ...interface{}) (int, error)
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB               UserDatabase
	PasswordStrength int
}

func (a *UserAPI) CreateUser(ctx *gin.Context) {
	user := model.UserRegisterRequest{}
	if err := ctx.Bind(&user); err == nil {
		user := a.toUser(&user, []byte{})
		existingUser, err := a.DB.GetUserByName(user.Name)
		if success := successOrAbort(ctx, 500, err); !success {
			return
		}
		if existingUser == nil {
			if success := successOrAbort(ctx, 500, a.DB.CreateUser(user)); !success {
				return
			}
			var emptyJson struct{}
			ctx.JSON(200, emptyJson)
		} else {
			ctx.AbortWithError(400, errors.New("This user is already exists"))
		}
	}
}

func (a *UserAPI) toUser(request *model.UserRegisterRequest, pw []byte) *model.User {
	user := &model.User{
		Name:  request.Name,
		Email: request.Email,
	}

	if request.Pass != "" {
		user.Pass = password.CreatePassword(request.Pass, a.PasswordStrength)
	} else {
		user.Pass = pw
	}

	return user
}
