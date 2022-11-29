package controller

import (
	"github.com/gin-gonic/gin"
	"go_app/helper"
	"go_app/service"
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

// Profile example
//
//	@Summary		Profile
//	@Produce		json
//	@Tags      User
//
// @Success		200		{string}	string			"ok"
// @Router			/user/profile [get]
// @Security ApiKeyAuth
func (c *userController) Profile(ctx *gin.Context) {

	id := c.jwtService.GetUserId(ctx)
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "ok!", user)
	ctx.JSON(http.StatusOK, res)
}
