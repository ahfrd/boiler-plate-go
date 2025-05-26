package controller

import (
	"encoding/json"
	"esb-code-assesment/app/model/request"
	"esb-code-assesment/app/model/response"
	"esb-code-assesment/app/service"
	"esb-code-assesment/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) CustomerController {
	return CustomerController{CustomerService: *customerService}
}

func (c *CustomerController) GetListCustomer(ctx *gin.Context) {
	requestId := guuid.New()

	bodyReq := &request.GetListCustomerRequest{
		PageNumber: ctx.Query("pageNumber"),
		PageSize:   ctx.Query("pageSize"),
	}
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.FailedGettingDataErrorBindingRequest.Code(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorBindingRequest.Message(), err),
			Status:  response.FailedGettingDataErrorBindingRequest.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	response, err := c.CustomerService.GetListCustomer(ctx, bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (c *CustomerController) AddCustomer(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.AddCustomerRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.CustomerService.AddCustomer(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
