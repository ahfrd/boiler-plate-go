package mock

import (
	"boiler-plate-rest/app/model/request"
	"boiler-plate-rest/app/model/response"

	"github.com/gin-gonic/gin"
)

type MockCustomerRepository struct {
	SelectDataCustomerFn func(ctx *gin.Context, request *request.GetListCustomerRequest) ([]response.CustomerInfoEntity, error)
	CountDataCustomerFn  func(ctx *gin.Context) (string, error)
	InsertCustomerFn     func(ctx *gin.Context, requst *request.AddCustomerRequest) error
}

func (ms *MockCustomerRepository) SelectDataCustomer(ctx *gin.Context, request *request.GetListCustomerRequest) ([]response.CustomerInfoEntity, error) {
	return ms.SelectDataCustomerFn(ctx, request)
}
func (ms *MockCustomerRepository) CountDataCustomer(ctx *gin.Context) (string, error) {
	return ms.CountDataCustomerFn(ctx)
}

func (ms *MockCustomerRepository) InsertCustomer(ctx *gin.Context, request *request.AddCustomerRequest) error {
	return ms.InsertCustomerFn(ctx, request)
}
