package service

import (
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/entity/response"
	"asia-quest/helpers"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	AuthRepository entity.AuthRepository
}

func NewAuthService(authRepository *entity.AuthRepository) entity.AuthService {
	return &AuthService{
		AuthRepository: *authRepository,
	}
}

func (s *AuthService) Login(ctx *gin.Context, params *request.LoginRequest, uid string) (*response.GeneralResponse, error) {
	combine := strings.ToUpper(params.Username) + params.Password
	hash := []byte(combine)
	hash_byte := sha256.Sum256(hash)
	hash_str := hex.EncodeToString(hash_byte[:])
	selectUsername, err := s.AuthRepository.Login(params)
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	fmt.Println("13123;213;'12;3';")
	fmt.Println(selectUsername.Password)
	if selectUsername.Password != hash_str {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "user/password salah",
		}, nil
	}

	layoutFormat := "2006-01-02 15:04:05"
	strtime, _ := time.Parse(layoutFormat, time.Now().String())
	sessionMaxTime := strtime.Add(time.Minute * 30)

	claims := &request.Claims{
		Username: selectUsername.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: sessionMaxTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(helpers.JwtKey())
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	http.SetCookie(ctx.Writer,
		&http.Cookie{
			Name:    "token",
			Path:    "/",
			Value:   tokenString,
			Expires: sessionMaxTime,
		})
	fmt.Println(tokenString)

	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
	}, nil
}
