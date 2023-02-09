package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	ValidateGoogleJWT(tokenString string) (GoogleClaims, error)
}

type Miner interface {
	GetDevice(address string, isIP bool) (app.MinerDevice, error)
	GetAllDevices() ([]app.MinerDevice, error)
	AddNew(dev app.MinerDevice) error
	AddDevices(model string, isIp bool, connections []string, locInfo [][]string) error
	GetDevicesByType(miner_type string) ([]app.MinerDevice, error)
	GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error)
	GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error)
	IsDeviceAdded(address string, isIP bool) (bool, error)
	IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error)
	GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error)
}

type User interface {
	GetUserByID(id int) (app.User, error)
}

type Service struct {
	Authorization
	Miner
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Miner:         NewMinerService(repos.Miner),
	}
}
