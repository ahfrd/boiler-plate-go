package mock

import (
	"boiler-plate-rest/app/model/request"
	"boiler-plate-rest/app/model/response"

	"github.com/gin-gonic/gin"
)

type MockCustomerController struct {
	GetListCustomerFn func(ctx *gin.Context, request *request.GetListCustomerRequest, uid string) (*response.GenericResponse, error)
	AddCustomerFn     func(ctx *gin.Context, request *request.AddCustomerRequest, uid string) (*response.GenericResponse, error)
}

func (m *MockCustomerController) GetListCustomer(ctx *gin.Context, request *request.GetListCustomerRequest, uid string) (*response.GenericResponse, error) {
	if m.GetListCustomerFn != nil {
		return m.GetListCustomerFn(ctx, request, uid)
	}
	return nil, nil
}

func (m *MockCustomerController) AddCustomer(ctx *gin.Context, request *request.AddCustomerRequest, uid string) (*response.GenericResponse, error) {
	if m.AddCustomerFn != nil {
		return m.AddCustomerFn(ctx, request, uid)
	}
	return nil, nil
}
