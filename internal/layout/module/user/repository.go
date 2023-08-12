package user

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/db"
)

type Repo struct {
	db *db.DbClient
}

func NewRepo(d *db.DbClient) *Repo{
	r:= new(Repo)
	r.db=d
	return r

}
