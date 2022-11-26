package controller

import (
	"github.com/gin-gonic/gin"
	"go_app/dto"
	"go_app/entity"
	"go_app/helper"
	"go_app/service"
	"net/http"
	"strconv"
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

// Login example
//
//	@Summary		Login
//	@Produce		json
//	@Tags      Auth
//
// @Param			email	query		string			true	"Email"
// @Param			password query		string			true	"Password"
// @Success		200		{string}	string			"ok"
// @Router			/auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to response request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VertifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "ok", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Pleace chek again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Upload example

// Register example
//
//	@Summary		Register
//	@Produce		json
//	@Tags      Auth
//
// @Param			name	query		string			true	"Name"
// @Param			email	query		string			true	"Email"
// @Param			password query		string			true	"Password"
// @Success		200		{string}	string			"ok"
// @Router			/auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "ok!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
