package controller

import (
	"go_study/task4/service"

	"github.com/gin-gonic/gin"
)

var LoginCtl *LoginController

type LoginController struct {
	userService *service.UserService
}

func NewLoginController() *LoginController {
	LoginCtl = &LoginController{
		userService: service.UserSve,
	}
	return LoginCtl
}

func (ctl *LoginController) Login(c *gin.Context) {
	ctl.userService.Login(c)
}
