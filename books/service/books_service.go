package service

import (
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/entity/response"
	"asia-quest/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type BooksService struct {
	BooksRepository entity.BooksRepository
}

func NewBooksService(booksRepository *entity.BooksRepository) entity.BooksService {
	return &BooksService{
		BooksRepository: *booksRepository,
	}
}

func (s *BooksService) Create(ctx *gin.Context, params *request.CreateRequest, uid string) (*response.GeneralResponse, error) {
	tokenJwt, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {

			return &response.GeneralResponse{
				Code: "400",
				Msg:  "no cookie found",
			}, nil
		}

		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}

	tokenStr := tokenJwt
	claims := &request.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": constant.GeneralErrorCode,
		// 	"msg":  constant.GeneralErrorDesc,
		// })
		// return
	}

	if !tkn.Valid {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "token not valid",
		}, nil
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": constant.RCTokenNotValid,
		// 	"msg":  "Token Not Valid",
		// })
		// return
	}
	params.Price = helpers.FormatIdr(params.Price)
	createUser := s.BooksRepository.Create(params)
	if createUser != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  createUser.Error(),
		}, nil
	}
	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
		Data: params,
	}, nil
}

func (s *BooksService) Read(ctx *gin.Context, params *request.ReadRequest, uid string) (*response.GeneralResponse, error) {
	tokenJwt, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {

			return &response.GeneralResponse{
				Code: "400",
				Msg:  "no cookie found",
			}, nil
		}

		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}

	tokenStr := tokenJwt
	claims := &request.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": constant.GeneralErrorCode,
		// 	"msg":  constant.GeneralErrorDesc,
		// })
		// return
	}

	if !tkn.Valid {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "token not valid",
		}, nil
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": constant.RCTokenNotValid,
		// 	"msg":  "Token Not Valid",
		// })
		// return
	}
	getBooks, err := s.BooksRepository.Read(params)
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	var resData response.ReadResponse
	resData.Title = getBooks.Title
	resData.Description = getBooks.Description
	category := strings.Split(getBooks.Category, ",")
	// Populate category with data

	var categoryBooks []response.CategoryBook
	for _, category := range category {
		categoryBook := response.CategoryBook{
			Data: category,
			// Additional fields initialization if required
		}
		categoryBooks = append(categoryBooks, categoryBook)
	}

	resData.Category = categoryBooks
	keyword := strings.Split(getBooks.Keyword, ",")
	var keywordBooks []response.KeywordBook
	for _, keyword := range keyword {
		KeywordBook := response.KeywordBook{
			Data: keyword,
		}
		keywordBooks = append(keywordBooks, KeywordBook)
	}
	resData.Keyword = keywordBooks
	resData.Price = getBooks.Price
	resData.Publisher = getBooks.Publisher

	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
		Data: resData,
	}, nil
}
func (s *BooksService) Update(ctx *gin.Context, params *request.UpdateRequest, uid string) (*response.GeneralResponse, error) {
	tokenJwt, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {

			return &response.GeneralResponse{
				Code: "400",
				Msg:  "no cookie found",
			}, nil
		}

		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}

	tokenStr := tokenJwt
	claims := &request.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": constant.GeneralErrorCode,
		// 	"msg":  constant.GeneralErrorDesc,
		// })
		// return
	}

	if !tkn.Valid {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "token not valid",
		}, nil
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": constant.RCTokenNotValid,
		// 	"msg":  "Token Not Valid",
		// })
		// return
	}

	params.Price = helpers.FormatIdr(params.Price)

	updateBooks := s.BooksRepository.Update(params)
	if updateBooks != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  updateBooks.Error(),
		}, nil
	}

	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
	}, nil
}

func (s *BooksService) Delete(ctx *gin.Context, params *request.DeleteRequest, uid string) (*response.GeneralResponse, error) {
	tokenJwt, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {

			return &response.GeneralResponse{
				Code: "400",
				Msg:  "no cookie found",
			}, nil
		}

		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}

	tokenStr := tokenJwt
	claims := &request.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": constant.GeneralErrorCode,
		// 	"msg":  constant.GeneralErrorDesc,
		// })
		// return
	}

	if !tkn.Valid {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "token not valid",
		}, nil
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": constant.RCTokenNotValid,
		// 	"msg":  "Token Not Valid",
		// })
		// return
	}
	delete := s.BooksRepository.Delete(params)
	if delete != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  delete.Error(),
		}, nil
	}
	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
	}, nil
}
