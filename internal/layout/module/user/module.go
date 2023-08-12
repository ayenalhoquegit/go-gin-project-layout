package user

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/db"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository *db.DbClient
}

func NewModule(db *db.DbClient) *Module {
	m := new(Module)
	m.Repository = NewRepo(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service)
	return m
}
