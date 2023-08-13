package user

import "github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/entity"

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
