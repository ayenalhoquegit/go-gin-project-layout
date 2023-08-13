package user

import (
	"database/sql"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/entity"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(d *sql.DB) *Repo {
	r := new(Repo)
	r.db = d
	return r

}

func (r *Repo) findAll() ([]entity.User, error) {
	var user []entity.User
	//rows, err := r.db.
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u entity.User
		rows.Scan(&u.Id, &u.Name, &u.Email, &u.Gender)
		user = append(user, u)
	}
	return user, nil
	//c.IndentedJSON(http.StatusOK, user)
}
