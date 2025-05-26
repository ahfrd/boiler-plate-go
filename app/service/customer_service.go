package service

import (
	"boiler-plate-rest/app/model/request"
	"boiler-plate-rest/app/model/response"
	"boiler-plate-rest/app/repository"
	"boiler-plate-rest/helpers"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	customerServiceImpl struct {
		CustomerRepository repository.CustomerRepository
	}
	CustomerService interface {
		GetListCustomer(ctx *gin.Context, request *request.GetListCustomerRequest, uid string) (*response.GenericResponse, error)
		AddCustomer(ctx *gin.Context, request *request.AddCustomerRequest, uid string) (*response.GenericResponse, error)
	}
)

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{
		CustomerRepository: customerRepository,
	}
}

func (s *customerServiceImpl) GetListCustomer(ctx *gin.Context, request *request.GetListCustomerRequest, uid string) (*response.GenericResponse, error) {
	var (
		countDataResponse response.CountData
		listDataInvoice   response.ListCustomerResponse

		res response.GenericResponse
	)
	countData, err := s.CustomerRepository.CountDataCustomer(ctx)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}
	if countData == "0" {
		errMsg := errors.New("data customer tidak dCustomerukan")
		helpers.LogError(ctx, errMsg.Error(), uid)

		return &response.GenericResponse{
			Code:    response.BadRequestNoDataFound.Code(),
			Status:  response.BadRequestNoDataFound.Status(),
			Message: fmt.Sprintf(response.BadRequestNoDataFound.Message(), errMsg.Error()),
		}, nil
	}

	paginationResponse := helpers.PaginationHelpers(countData, request.PageNumber, request.PageSize)
	marshalPaging, _ := json.Marshal(paginationResponse)
	if err := json.Unmarshal(marshalPaging, &countDataResponse); err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Status:  response.BadRequestMarshalError.Status(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err.Error()),
		}, nil
	}
	request.PageNumber = countDataResponse.FirstRecord

	selectDataCustomer, err := s.CustomerRepository.SelectDataCustomer(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}
	listDataInvoice.ListData = selectDataCustomer
	listDataInvoice.Pagination = countDataResponse

	res.Status = response.Success.Status()
	res.Code = response.Success.Code()
	res.Message = fmt.Sprintf(response.Success.Message(), " get all customer")
	res.Data = listDataInvoice
	return &res, nil
}

func (s *customerServiceImpl) AddCustomer(ctx *gin.Context, request *request.AddCustomerRequest, uid string) (*response.GenericResponse, error) {
	var res response.GenericResponse

	if err := s.CustomerRepository.InsertCustomer(ctx, request); err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err),
		}, nil

	}
	res.Status = response.Success.Status()
	res.Code = response.Success.Code()
	res.Message = fmt.Sprintf(response.Success.Message(), " create Customer")
	return &res, nil
}
