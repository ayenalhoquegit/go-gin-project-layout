package user

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/dto"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/entity"
)

type Service struct {
	repository *Repo
}

func NewService(r *Repo) *Service {
	s := new(Service)
	s.repository = r
	return s
}
func (s *Service) findAll() ([]entity.User, error) {
	user, err := s.repository.findAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) CreateUser(d *dto.UserDto) (entity.User, error) {
	u := entity.User{}
	u.Name = d.Name
	u.Email = d.Email
	u.Gender = d.Gender
	err := s.repository.CreateUser(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) FindUser(id int) (entity.User, error) {
	u, err := s.repository.FindUser(id)
	if err != nil {
		return u, err
	}
	return u, nil
}
func (s *Service) DeleteUser(id int) (int64, error) {
	affect, err := s.repository.DeleteUser(id)
	if err != nil {
		return affect, err
	}
	return affect, nil
}
