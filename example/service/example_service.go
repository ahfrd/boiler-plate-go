package service

import (
	"example-boiler-plate/entity"
	"example-boiler-plate/entity/request"
	"example-boiler-plate/entity/response"

	"github.com/gin-gonic/gin"
)

type ExampleService struct {
	AuthRepository entity.ExampleRepository
}

func NewExampleService(authRepository *entity.ExampleRepository) entity.ExampleService {
	return &ExampleService{
		AuthRepository: *authRepository,
	}
}

func (s *ExampleService) ExampleService(ctx *gin.Context, params *request.ExampleRequest, uid string) (*response.LoginResponse, error) {
	return nil, nil
}
