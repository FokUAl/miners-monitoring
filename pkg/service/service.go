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

type Miner interface {
	GetDevice(ip_address string) (app.MinerDevice, error)
	GetAllDevices() ([]app.MinerDevice, error)
	AddNew(dev app.MinerDevice) error
	AddDevices(model string, isIp bool, connections, shelfData, rowData, columnData []string) error
}

type User interface {
	GetUser(username, password string) (app.User, error)
}

type Service struct {
	Authorization
	Miner
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Miner:         NewMinerService(repos.Miner),
	}
}
