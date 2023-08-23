package auth

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
)

type Module struct {
	Service *AuthService
	Handler *Handler
}

func NewModule(s *user.Service) *Module {
	m := new(Module)
	m.Service = NewService(s)
	m.Handler = NewHandler(s)
	return m
}
