package routes

import (
	"asia-quest/auth/controller"

	"github.com/gin-gonic/gin"
)

func SetUpAuthRoute(router *gin.Engine, booksController *controller.AuthController) {
	router.POST("/login", booksController.Login)
	router.GET("/oauth-login", booksController.OauthLogin)
	router.GET("/callback", booksController.OauthCallback)
	router.GET("/", booksController.PageMain)
}
