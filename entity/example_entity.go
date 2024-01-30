package entity

import (
	"example-boiler-plate/entity/request"
	"example-boiler-plate/entity/response"

	"github.com/gin-gonic/gin"
)

type ExampleService interface {
	ExampleService(ctx *gin.Context, request *request.ExampleRequest, uid string) (*response.LoginResponse, error)
}

type ExampleRepository interface {
	ExampleRepository(request *request.ExampleRequest) (*response.LoginResponse, error)
}
