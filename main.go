package main

import (
	authController "asia-quest/auth/controller"
	authRepository "asia-quest/auth/repository"
	authService "asia-quest/auth/service"
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
	authRepository := authRepository.NewAuthRepository()
	//Service
	authService := authService.NewAuthService(&authRepository)
	booksService := booksService.NewBooksService(&booksRepository)
	//Controller
	booksController := booksController.NewBooksController(&booksService)
	authController := authController.NewAuthController(&authService)
	routes.SetUpRoutesBooks(router, &booksController)
	routes.SetUpAuthRoute(router, &authController)
	router.Run(":8000")
}
