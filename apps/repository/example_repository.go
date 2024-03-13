package repository

import (
	"example-boiler-plate/apps/entity"
	"example-boiler-plate/apps/entity/request"
	"example-boiler-plate/apps/entity/response"
)

type exampleRepository struct {
}

func NewExampleRepository() entity.ExampleRepository {
	return &exampleRepository{}
}
func (r *exampleRepository) ExampleRepository(request *request.ExampleRequest) (*response.LoginResponse, error) {

	return nil, nil
}
