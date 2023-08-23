package auth

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/entity"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/httppkg"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/jwtpkg"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Service *user.Service
}

func NewService(s *user.Service) *AuthService {
	m := new(AuthService)
	m.Service = s
	return m
}

func (m *AuthService) Authorize(ctx *gin.Context) (entity.User, error) {
	var u entity.User
	splits, err := httppkg.ParseAuthToken(ctx)
	if err != nil {
		return u, err
	}
	tokenBody := splits[1]
	claims, err := jwtpkg.VerifyToken(tokenBody)
	if err != nil {
		return u, err
	}
	// find user
	u, errr := m.Service.FindUser(claims.Payload.Id)
	if errr != nil {
		return u, err
	}
	return u, nil
}
