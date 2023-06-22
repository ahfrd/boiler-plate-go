package routes

import (
	"asia-quest/auth/controller"

	"github.com/gin-gonic/gin"
)

func SetUpAuthRoute(router *gin.Engine, booksController *controller.AuthController) {
	v1 := router.Group("/api/v1/auth")
	{
		v1.POST("/login", booksController.Login)
	}
}
