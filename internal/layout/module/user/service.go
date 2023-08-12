package user

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/db"
)

type Service struct{
	repository  *db.DbClient
}

func NewService(r *db.DbClient) *Service {
	s := new(Service)
	s.repository = r
	return s
}