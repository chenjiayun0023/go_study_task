package controller

import (
	"go_study/task4/service"

	"github.com/gin-gonic/gin"
)

var UserCtl *UserController

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	UserCtl = &UserController{
		userService: service.UserSve,
	}
	return UserCtl
}

func (ctl *UserController) Register(c *gin.Context) {
	ctl.userService.Register(c)
}
