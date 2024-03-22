package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/gokominfo-unittest/internal/handler"
	"github.com/rifqidamarali/gokominfo-unittest/internal/middleware"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	// activity
	// /users/sign-up
	u.v.POST("/sign-up", u.handler.UserSignUp)

	// users
	u.v.Use(middleware.CheckAuthBearer)
	// /users
	u.v.GET("", u.handler.GetUsers)
	// /users/:id
	u.v.GET("/:id", u.handler.GetUsersById)
	u.v.DELETE("/:id", u.handler.DeleteUsersById)
}
