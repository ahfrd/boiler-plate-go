package controller_test

import (
	"boiler-plate-rest/app/controller"
	"boiler-plate-rest/app/model/mock"
	"boiler-plate-rest/app/model/request"
	"boiler-plate-rest/app/model/response"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetListCustomer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.CustomerController{}

	router := gin.Default()
	router.GET("/customer/list-customer", controller.GetListCustomer)

	t.Run("Successful", func(t *testing.T) {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/customer/list-customer?pageNumber=1&pageSize=10", nil)

		controller.CustomerService = &mock.MockCustomerController{
			GetListCustomerFn: func(ctx *gin.Context, req *request.GetListCustomerRequest, requestId string) (*response.GenericResponse, error) {
				response := &response.GenericResponse{
					Code:    response.Success.Code(),
					Message: fmt.Sprintf(response.Success.Message(), "sukses"),
					Status:  response.Success.Status(),
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Customer - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/customer/list-customer", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Customer - Service Error", func(t *testing.T) {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/customer/list-customer?pageNumber=1&pageSize=10", nil)

		controller.CustomerService = &mock.MockCustomerController{
			GetListCustomerFn: func(ctx *gin.Context, req *request.GetListCustomerRequest, requestId string) (*response.GenericResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}

func TestAddCustomer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.CustomerController{}

	router := gin.Default()
	router.POST("/customer/add-customer", controller.AddCustomer)

	t.Run("Successful", func(t *testing.T) {
		bodyReq := request.AddCustomerRequest{
			CustomerName: "Metro Cafe",
			Address:      "Jl. LLLLLLLL",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/customer/add-customer", bytes.NewReader(requestData))

		controller.CustomerService = &mock.MockCustomerController{
			AddCustomerFn: func(ctx *gin.Context, req *request.AddCustomerRequest, requestId string) (*response.GenericResponse, error) {
				response := &response.GenericResponse{
					Code:    response.Success.Code(),
					Message: fmt.Sprintf(response.Success.Message(), "sukses"),
					Status:  response.Success.Status(),
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Customer - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/customer/add-customer", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Customer - Service Error", func(t *testing.T) {
		bodyReq := request.AddCustomerRequest{
			CustomerName: "Metro Cafe",
			Address:      "Jl. LLLLLLLL",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/customer/add-customer", bytes.NewReader(requestData))

		controller.CustomerService = &mock.MockCustomerController{
			AddCustomerFn: func(ctx *gin.Context, req *request.AddCustomerRequest, requestId string) (*response.GenericResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}
