package controller

import (
	"github.com/gin-gonic/gin"
	"golang_app/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		jwtService:  jwtService,
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
}

func (c *authController) Register(ctx *gin.Context) {
}
