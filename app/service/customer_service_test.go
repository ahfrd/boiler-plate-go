package service_test

import (
	"errors"
	mockModel "esb-code-assesment/app/model/mock"
	"esb-code-assesment/app/model/request"
	"esb-code-assesment/app/model/response"
	"esb-code-assesment/app/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomer(t *testing.T) {

	mockRepo := &mockModel.MockCustomerRepository{}
	svc := service.NewCustomerService(mockRepo)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/customer/list-customer", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	t.Run("Success", func(t *testing.T) {

		mockRepo.CountDataCustomerFn = func(ctx *gin.Context) (string, error) {
			return "2", nil
		}
		mockRepo.SelectDataCustomerFn = func(ctx *gin.Context, req *request.GetListCustomerRequest) ([]response.CustomerInfoEntity, error) {
			return []response.CustomerInfoEntity{
				{CustomerName: "0000015"},
				{CustomerName: "0000016"},
			}, nil
		}

		req := &request.GetListCustomerRequest{
			PageNumber: "1",
			PageSize:   "10",
		}

		resp, err := svc.GetListCustomer(ctx, req, "uid")

		assert.NoError(t, err)
		assert.Equal(t, response.Success.Code(), resp.Code)
		assert.Equal(t, "success  get all customer", resp.Message)

	})

}

func TestAddCustomer(t *testing.T) {

	mockRepo := &mockModel.MockCustomerRepository{}
	svc := service.NewCustomerService(mockRepo)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/customer/add-customer", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	t.Run("Success", func(t *testing.T) {

		mockRepo.InsertCustomerFn = func(ctx *gin.Context, req *request.AddCustomerRequest) error {
			return nil
		}

		req := &request.AddCustomerRequest{
			CustomerName: "Metro Cafe",
			Address:      "Jl. LLLLLLLL",
		}

		resp, err := svc.AddCustomer(ctx, req, "uid")

		assert.NoError(t, err)
		assert.Equal(t, response.Success.Code(), resp.Code)
		assert.Equal(t, "success  create Customer", resp.Message)

	})

	t.Run("Failed insert Customer", func(t *testing.T) {

		mockRepo.InsertCustomerFn = func(ctx *gin.Context, req *request.AddCustomerRequest) error {
			return errors.New("failed")
		}

		req := &request.AddCustomerRequest{
			CustomerName: "Metro Cafe",
			Address:      "Jl. LLLLLLLL",
		}

		resp, err := svc.AddCustomer(ctx, req, "uid")

		assert.NoError(t, err)
		assert.Equal(t, response.FailedCreatingDataErrorInternal.Code(), resp.Code)
		assert.Equal(t, "error creating data failed", resp.Message)

	})

}
