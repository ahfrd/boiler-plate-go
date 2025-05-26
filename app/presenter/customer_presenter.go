package presenter

import (
	"esb-code-assesment/app/controller"
	"esb-code-assesment/app/repository"
	"esb-code-assesment/app/service"
	"esb-code-assesment/env"

	"github.com/gin-gonic/gin"
)

func mainCustomerPresenter(di *env.Dependency) *controller.CustomerController {
	customerRepository := repository.NewCustomerRepository(di)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(&customerService)
	return &customerController
}

func GetListCustomer(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainCustomerPresenter(env).GetListCustomer(ctx)
	}
}

func AddCustomer(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainCustomerPresenter(env).AddCustomer(ctx)
	}
}
