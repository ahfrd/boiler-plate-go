package main

import (
	booksController "asia-quest/books/controller"
	booksRepository "asia-quest/books/repository"
	booksService "asia-quest/books/service"
	"asia-quest/helpers"
	"asia-quest/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	router := gin.Default()
	booksRepository := booksRepository.NewBooksRepository()
	//Service
	booksService := booksService.NewBooksService(&booksRepository)
	//Controller
	booksController := booksController.NewBooksController(&booksService)
	routes.SetUpRoutesBooks(router, &booksController)
	router.Run(":8000")
}
