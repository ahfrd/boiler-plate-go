package routes

import (
	"asia-quest/books/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoutesBooks(router *gin.Engine, booksController *controller.BooksController) {
	v1 := router.Group("/api/v1/crud")
	{
		v1.POST("/create", booksController.Create)
		v1.POST("/read", booksController.Read)
		v1.POST("/update", booksController.Update)
		v1.POST("/delete", booksController.Delete)

	}
}
