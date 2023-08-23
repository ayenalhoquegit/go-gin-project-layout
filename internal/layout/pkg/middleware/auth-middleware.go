package middleware

import (
	"fmt"
	"net/http"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/auth"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Service *auth.AuthService
}

func NewAuthMiddleware(s *auth.AuthService) *AuthMiddleware {
	m := new(AuthMiddleware)
	m.Service = s
	return m
}

func (m *AuthMiddleware) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Fire middleware")
		e, err := m.Service.Authorize(c)
		if err != nil {
			// send error response
			response.RespondError(http.StatusForbidden, err, c)
			// need to abort the middleware chain
			c.Abort()
			return
		}
		fmt.Println("User : ", e)
		c.Set("user", e)

		c.Next()
	}
}
