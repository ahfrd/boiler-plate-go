package repository

import (
	"example-boiler-plate/entity"
	"example-boiler-plate/entity/request"
	"example-boiler-plate/entity/response"
)

type exampleRepository struct {
}

func NewExampleRepository() entity.ExampleRepository {
	return &exampleRepository{}
}
func (r *exampleRepository) ExampleRepository(request *request.ExampleRequest) (*response.LoginResponse, error) {

	return nil, nil
}
