package controller

import(
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}




func Login(ctx *gin.Context) {
}

func Register(ctx *gin.Context) {
}