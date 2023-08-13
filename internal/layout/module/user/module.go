package user

import (
	"database/sql"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository *Repo
}

func NewModule(db *sql.DB) *Module {
	m := new(Module)
	m.Repository = NewRepo(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service)
	return m
}
