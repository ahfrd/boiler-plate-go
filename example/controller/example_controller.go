package controller

import (
	"encoding/json"
	"example-boiler-plate/entity"
	"example-boiler-plate/entity/request"
	"example-boiler-plate/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type ExampleController struct {
	ExampleService entity.ExampleService
}

func NewExampleController(exampleService *entity.ExampleService) ExampleController {
	return ExampleController{ExampleService: *exampleService}
}

func (c *ExampleController) Login(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.ExampleRequest
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

	response, err := c.ExampleService.ExampleService(ctx, &bodyReq, requestId.String())
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
