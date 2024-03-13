package routes

import (
	"example-boiler-plate/apps/controller"

	"github.com/gin-gonic/gin"
)

func SetUpExampleRoute(router *gin.Engine, exampleController *controller.ExampleController) {
	router.POST("/login", exampleController.Login)

}
