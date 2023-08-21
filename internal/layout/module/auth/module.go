package auth

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
)

type Module struct {
	Service *user.Service
	Handler *Handler
}

func NewModule(service *user.Service) *Module {
	m := new(Module)
	m.Service = service
	m.Handler = NewHandler(m.Service)
	return m
}
