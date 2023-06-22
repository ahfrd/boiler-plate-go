package entity

import (
	"asia-quest/entity/request"
	"asia-quest/entity/response"

	"github.com/gin-gonic/gin"
)

type BooksService interface {
	Create(ctx *gin.Context, request *request.CreateRequest, uid string) (*response.GeneralResponse, error)
	Read(ctx *gin.Context, request *request.ReadRequest, uid string) (*response.GeneralResponse, error)
	Update(ctx *gin.Context, request *request.UpdateRequest, uid string) (*response.GeneralResponse, error)
	Delete(ctx *gin.Context, request *request.DeleteRequest, uid string) (*response.GeneralResponse, error)
}

type BooksRepository interface {
	Create(request *request.CreateRequest) error
	Read(request *request.ReadRequest) (*response.ReadQueryEntity, error)
	Update(request *request.UpdateRequest) error
	Delete(request *request.DeleteRequest) error
}
