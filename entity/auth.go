package entity

import (
	"asia-quest/entity/request"
	"asia-quest/entity/response"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context, request *request.LoginRequest, uid string) (*response.GeneralResponse, error)
}

type AuthRepository interface {
	Login(request *request.LoginRequest) (*response.LoginResponse, error)
}
