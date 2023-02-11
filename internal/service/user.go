package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(id int) (app.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetRole(username string) (string, error) {
	return s.repo.GetUserRole(username)
}
