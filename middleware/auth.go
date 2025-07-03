package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const ClaimsContextKey = "jwt_claims"

var secretKeyJWT = "ITS_SECRET"

// Your custom claims
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Generic JWT service
type JWTService[T jwt.Claims] struct {
	Secret    []byte
	NewClaims func() T
}

func NewJWTMiddleware() gin.HandlerFunc {
	svc := &JWTService[*MyClaims]{
		Secret: []byte(secretKeyJWT),
		NewClaims: func() *MyClaims {
			return &MyClaims{}
		},
	}
	return JWTMiddleware(svc)
}

// Generic JWT middleware
func JWTMiddleware[T jwt.Claims](svc *JWTService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := svc.Validate(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(claims)
		c.Set(ClaimsContextKey, claims)
		c.Next()
	}
}

func (s *JWTService[T]) Validate(tokenStr string) (T, error) {
	var zero T

	claims := s.NewClaims()

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return s.Secret, nil
	})
	if err != nil || !token.Valid {
		return zero, err
	}

	expire, err := claims.GetExpirationTime()
	if err != nil {
		return zero, err
	}
	if expire.Before(time.Now()) {
		return zero, errors.New("token expire")
	}

	return token.Claims.(T), nil
}

func GenerateJWT[T jwt.Claims](claims T) (string, error) {
	if c, ok := any(claims).(interface{ setDefaults() }); ok {
		c.setDefaults()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKeyJWT))
}

func (c *MyClaims) setDefaults() {
	now := time.Now()
	c.IssuedAt = jwt.NewNumericDate(now)
	c.ExpiresAt = jwt.NewNumericDate(now.Add(time.Hour))
	c.ID = uuid.NewString() // jti
}
