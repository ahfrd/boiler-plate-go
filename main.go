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
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	fmt.Println(os.Getenv("CLIENT_ID"))
	fmt.Println(os.Getenv("REDIRECT"))
	oauthConfig := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	router := gin.Default()
	booksRepository := booksRepository.NewBooksRepository()
	authRepository := authRepository.NewAuthRepository()
	//Service
	authService := authService.NewAuthService(&authRepository, &oauthConfig)
	booksService := booksService.NewBooksService(&booksRepository)
	//Controller
	booksController := booksController.NewBooksController(&booksService)
	authController := authController.NewAuthController(&authService)
	routes.SetUpRoutesBooks(router, &booksController)
	routes.SetUpAuthRoute(router, &authController)
	router.Run(":8080")
}
