package user

import (
	"database/sql"
	"log"

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

func (r *Repo) CreateUser(u *entity.User) error {
	result, err := r.db.Exec("INSERT INTO users (name, email, gender) VALUES (?, ?, ?)", u.Name, u.Email, u.Gender)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	return nil
}

func (r *Repo) FindUser(id int) (entity.User, error) {
	var user entity.User
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Gender)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repo) DeleteUser(id int) (int64, error) {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affect, nil

}
