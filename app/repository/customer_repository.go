package repository

import (
	"esb-code-assesment/app/model/request"
	"esb-code-assesment/app/model/response"
	"esb-code-assesment/env"
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	CustomerRepository interface {
		SelectDataCustomer(ctx *gin.Context, request *request.GetListCustomerRequest) ([]response.CustomerInfoEntity, error)
		CountDataCustomer(ctx *gin.Context) (string, error)
		InsertCustomer(ctx *gin.Context, request *request.AddCustomerRequest) error
	}
	customerRepositoryImpl struct {
		di *env.Dependency
	}
)

func NewCustomerRepository(di *env.Dependency) CustomerRepository {
	return &customerRepositoryImpl{
		di: di,
	}
}

func (r *customerRepositoryImpl) SelectDataCustomer(ctx *gin.Context, request *request.GetListCustomerRequest) ([]response.CustomerInfoEntity, error) {
	var results []response.CustomerInfoEntity
	q := `select customer_name,address from tbl_customer order by id asc limit ?,?`
	result, err := r.di.DB.QueryContext(ctx, q, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var dataCust response.CustomerInfoEntity
		if err = result.Scan(&dataCust.CustomerName, &dataCust.Address); err != nil {
			return nil, err
		}
		results = append(results, dataCust)

	}

	return results, nil
}

func (r *customerRepositoryImpl) CountDataCustomer(ctx *gin.Context) (string, error) {
	var count string
	q := `select count(*) as count from tbl_customer`
	if err := r.di.DB.QueryRowContext(ctx, q).Scan(&count); err != nil {
		return "", err
	}
	return count, nil

}

func (r *customerRepositoryImpl) InsertCustomer(ctx *gin.Context, request *request.AddCustomerRequest) error {
	q := `INSERT INTO tbl_customer (customer_name,address) values (?,?)`
	_, err := r.di.DB.ExecContext(ctx, q, request.CustomerName, request.Address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
