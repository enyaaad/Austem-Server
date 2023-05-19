package handlers

import (
	"AustemServer/package/auth"
	"github.com/gin-gonic/gin"
)

func SignInEndpoint(router *gin.RouterGroup, usecase auth.UseCase) {
	h := newHandler(usecase)

	router.POST("/", h.signIn)
}
