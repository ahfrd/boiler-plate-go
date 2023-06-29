package controller

import (
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type AuthController struct {
	AuthService entity.AuthService
}

func NewAuthController(authService *entity.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.LoginRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.AuthService.Login(ctx, &bodyReq, requestId.String())
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

func (c *AuthController) OauthLogin(ctx *gin.Context) {
	requestId := guuid.New()

	logStart := helpers.LogRequest(ctx, "Oauth Login", requestId.String())
	fmt.Println(logStart)

	response, err := c.AuthService.OauthLogin(ctx, requestId.String())
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
	ctx.Redirect(http.StatusTemporaryRedirect, response)
}
func (c *AuthController) OauthCallback(ctx *gin.Context) {
	requestId := guuid.New()

	logStart := helpers.LogRequest(ctx, "Oauth callback", requestId.String())
	fmt.Println(logStart)
	code := ctx.Query("code")
	state := ctx.Query("state")
	err := c.AuthService.OauthCallback(ctx, code, state)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, "Oauth callback", requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, "Sukses")
}

func (c *AuthController) PageMain(ctx *gin.Context) {
	const IndexPage = `
	<html>
		<head>
			<title>OAuth-2 Test</title>
		</head>
		<body>
			<h2>OAuth-2 Test</h2>
			<p>
				Login with the following,
			</p>
			<ul>
				<li><a href="/oauth-login">Google</a></li>
			</ul>
		</body>
	</html>
	`
	ctx.Header("Content-Type", "text/html, charset=utf-8")
	ctx.Status(http.StatusOK)
	ctx.Writer.Write([]byte(IndexPage))

}
