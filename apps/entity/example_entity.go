package entity

import (
	"example-boiler-plate/apps/entity/request"
	"example-boiler-plate/apps/entity/response"

	"github.com/gin-gonic/gin"
)

type ExampleService interface {
	ExampleService(ctx *gin.Context, request *request.ExampleRequest, uid string) (*response.LoginResponse, error)
}

type ExampleRepository interface {
	ExampleRepository(request *request.ExampleRequest) (*response.LoginResponse, error)
}
