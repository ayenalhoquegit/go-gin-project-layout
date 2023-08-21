package user

import (
	"net/http"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/dto"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/entity"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/constant"
	errorPkg "github.com/ayenalhoquegit/go-gin-project-layout/pkg/error"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repo
}

func NewService(r *Repo) *Service {
	s := new(Service)
	s.repository = r
	return s
}
func (s *Service) findAll() ([]entity.User, *errorPkg.HTTPErrorPkg) {
	user, err := s.repository.findAll()
	if err != nil {
		return user, errorPkg.HandleError(err)
	}
	return user, nil
}

func (s *Service) CreateUser(d *dto.UserDto) (entity.User, *errorPkg.HTTPErrorPkg) {
	u := entity.User{}
	u.Name = d.Name
	u.Email = d.Email
	u.Gender = d.Gender
	//u.Password = d.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	u.Password = string(hash)
	err = s.repository.CreateUser(&u)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	return u, nil
}

func (s *Service) FindUser(id int) (entity.User, *errorPkg.HTTPErrorPkg) {
	u, err := s.repository.FindUser(id)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	return u, nil
}
func (s *Service) FindUserByEmail(email string) (entity.User, *errorPkg.HTTPErrorPkg) {
	u, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	return u, nil
}

func (s *Service) UpdateUser(id int, dto *dto.UserDto) (entity.User, *errorPkg.HTTPErrorPkg) {
	u, err := s.repository.FindUser(id)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	u.Name = dto.Name
	u.Email = dto.Name
	u.Gender = dto.Gender
	rows, err := s.repository.UpdateUser(id, &u)
	if err != nil {
		return u, errorPkg.HandleError(err)
	}
	if rows > 0 {
		return u, nil
	}

	return u, &errorPkg.HTTPErrorPkg{Code: http.StatusBadRequest, Err: errorPkg.NewError(constant.OperationNotSuccess)}
}
func (s *Service) DeleteUser(id int) (int64, *errorPkg.HTTPErrorPkg) {
	affectedRows, err := s.repository.DeleteUser(id)
	if err != nil {
		return affectedRows, errorPkg.HandleError(err)
	}
	if affectedRows > 0 {
		return affectedRows, nil
	}
	return 0, &errorPkg.HTTPErrorPkg{Code: http.StatusBadRequest, Err: errorPkg.NewError(constant.OperationNotSuccess)}

}
