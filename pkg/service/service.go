package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	ValidateGoogleJWT(tokenString string) (GoogleClaims, error)
}

type Room interface{}

type Player interface {
	GetUser(username, password string) (app.User, error)
}

type Service struct {
	Authorization
	Room
	Player
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
