package service

import (
	"example-boiler-plate/apps/entity"
	"example-boiler-plate/apps/entity/request"
	"example-boiler-plate/apps/entity/response"

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
