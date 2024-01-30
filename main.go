package main

import (
	exampleController "example-boiler-plate/example/controller"
	exampleRepository "example-boiler-plate/example/repository"
	exampleService "example-boiler-plate/example/service"
	"example-boiler-plate/helpers"
	"example-boiler-plate/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	router := gin.Default()
	ExampleRepository := exampleRepository.NewExampleRepository()
	//Service
	ExampleService := exampleService.NewExampleService(&ExampleRepository)
	//Controller
	ExampleController := exampleController.NewExampleController(&ExampleService)
	routes.SetUpExampleRoute(router, &ExampleController)
	router.Run(":8080")
}
