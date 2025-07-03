package wire

import (
	"boiler-plate-rest/app/controller"
	"boiler-plate-rest/app/repository"
	"boiler-plate-rest/app/service"
	"boiler-plate-rest/env"

	"github.com/gin-gonic/gin"
)

func mainCustomerWire(di *env.Dependency) *controller.CustomerController {
	customerRepository := repository.NewCustomerRepository(di)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(&customerService)
	return &customerController
}

func GetListCustomer(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainCustomerWire(env).GetListCustomer(ctx)
	}
}

func AddCustomer(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainCustomerWire(env).AddCustomer(ctx)
	}
}

func TestingGenerateJWT(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainCustomerWire(env).TestingGenerateJWT(ctx)
	}
}
