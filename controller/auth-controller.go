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

// Upload example

// Login example
//
//	@Summary		TEST file
//	@Description	Test test
//	@Produce		json
//	@Tags      Test
//
// @Param			email	path		string			true	"Email"
// @Param			password path		string			true	"Password"
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

// GetStructArrayByString example
//
//	@Description	get struct array by ID
//	@ID				get-struct-array-by-string
//	@Accept			json
//	@Produce		json
//	@Param			some_id	path		string			true	"Some ID"
//	@Param			offset	query		int				true	"Offset"
//	@Param			limit	query		int				true	"Offset"
//	@Success		200		{string}	string			"ok"
//	@Router			/testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(w http.ResponseWriter, r *http.Request) {
	// write your code
}
func Upload(w http.ResponseWriter, r *http.Request) {
	// write your code
}

// AnonymousField example
//
//	@Summary	use Anonymous field
//	@Success		200		{string}	string			"ok"
func AnonymousField() {

}

// Pet3 example
type Pet3 struct {
	ID int `json:"id"`
}
