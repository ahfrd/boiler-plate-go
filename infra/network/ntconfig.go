package network

import (
	"bytes"
	"esb-code-assesment/app/model/response"
	"esb-code-assesment/app/presenter"
	"esb-code-assesment/env"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func InitRoutesGin(env *env.Dependency) *gin.Engine {
	rD := gin.New()
	rD.Use(CORSMiddleware())
	rD.Use(gin.CustomRecovery(recoverPanic))
	rC := rD.Group("api/admin/v1/")
	{
		customerGroup(rC.Group("/customer"), env)

	}
	return rD

}

func customerGroup(group *gin.RouterGroup, env *env.Dependency) {
	group.GET("/list-customer", presenter.GetListCustomer(env))
	group.POST("/add-customer", presenter.AddCustomer(env))

}

// recoverPanic : custom handler to recover from panic
func recoverPanic(c *gin.Context, _ interface{}) {
	var (
		msg   = "panic on %s line %v func %s"
		dot   = []byte(".")
		slash = []byte("/")
	)

	pc, _, line, ok := runtime.Caller(4)
	if !ok {
		msg = "server error"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		msg = "server error"
	}
	name := []byte(fn.Name())
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}

	msg = fmt.Sprintf(msg, fn.Name(), line, string(name))
	c.AbortWithStatusJSON(http.StatusInternalServerError,
		response.GenericResponse{
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), msg),
			Data:    nil,
		})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Client-ID, Client-Module, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
