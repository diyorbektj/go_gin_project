package controller

import (
	"github.com/gin-gonic/gin"
	"golang_app/service"
)

type UserController interface {
	Profile(ctx *gin.Context)
}
type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *userController) Profile(ctx *gin.Context) {

}
