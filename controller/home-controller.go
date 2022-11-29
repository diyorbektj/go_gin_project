package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type HomeController interface {
	Home(context *gin.Context)
}

type homeController struct{}

func NewHomeController() HomeController {
	return &homeController{}
}

func (h homeController) Home(ctx *gin.Context) {
	q := url.Values{}
	location := url.URL{Path: "/swagger/index.html", RawQuery: q.Encode()}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
