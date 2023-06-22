package request

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
