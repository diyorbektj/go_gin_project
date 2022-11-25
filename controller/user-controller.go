package controller

import (
	"github.com/gin-gonic/gin"
	"golang_app/helper"
	"golang_app/service"
	"net/http"
)

type UserController interface {
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Profile(ctx *gin.Context) {

	id := c.jwtService.GetUserId(ctx)
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "ok!", user)
	ctx.JSON(http.StatusOK, res)
}
