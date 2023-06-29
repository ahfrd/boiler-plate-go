package service

import (
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/entity/response"
	"asia-quest/helpers"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

type AuthService struct {
	OauthConfig    oauth2.Config
	AuthRepository entity.AuthRepository
}
type Email struct {
	Email string `json:"email"`
}

func NewAuthService(authRepository *entity.AuthRepository, oauthConfig *oauth2.Config) entity.AuthService {
	return &AuthService{
		AuthRepository: *authRepository,
		OauthConfig:    *oauthConfig,
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

	sessionMaxTime := time.Now().Add(time.Minute * 30)
	fmt.Println(sessionMaxTime)
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

func (s *AuthService) OauthLogin(ctx *gin.Context, uid string) (string, error) {

	url := s.OauthConfig.AuthCodeURL(uid)
	return url, nil
}
func (s *AuthService) OauthCallback(ctx *gin.Context, code string, state string) error {
	token, err := s.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return err
	}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return err
	}
	fmt.Println(response.Body)
	fmt.Println("././././")
	fmt.Println(response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var email Email
	err = json.Unmarshal(body, &email)
	if err != nil {
		return err
	}
	fmt.Println(email.Email)
	fmt.Println("==============")
	return nil
}
