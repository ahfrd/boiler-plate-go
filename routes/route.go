package routes

import (
	"example-boiler-plate/example/controller"

	"github.com/gin-gonic/gin"
)

func SetUpExampleRoute(router *gin.Engine, exampleController *controller.ExampleController) {
	router.POST("/login", exampleController.Login)

}
