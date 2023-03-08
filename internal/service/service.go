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
	GetDevice(address string) (app.MinerDevice, error)
	GetDevicesInfo() ([]app.AddInfo, error)
	GetAllDevices() ([]app.MinerDevice, error)
	AddNew(dev app.MinerDevice) error
	MappDevices(mappingInfo []app.AddInfo) error
	// GetDevicesByType(miner_type string) ([]app.MinerDevice, error)
	// GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error)
	// GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error)
	IsDeviceAdded(address string) (bool, error)
	IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error)
	GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error)
	GetDevicesByUser(username string) ([]app.MinerDevice, error)
	IsIPFree(ip_address string) (bool, error)
}

type User interface {
	GetUserByID(id int) (app.User, error)
	GetRole(username string) (string, error)
}

type Info interface {
	PingDevices() ([]string, error)
	SaveMinerData(data app.MinerData, ip_address string) error
	Transform(devices []app.MinerDevice) (map[string][]app.MinerData, error)
}

type Service struct {
	Authorization
	Miner
	User
	Info
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Miner:         NewMinerService(repos.Miner),
		Info:          NewInfoService(repos.Info),
	}
}
