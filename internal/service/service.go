package service

import (
	"time"

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
	IsDeviceAdded(address string) (bool, error)
	IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error)
	GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error)
	GetDevicesByUser(username string) ([]app.MinerDevice, error)
	IsAddressFree(address string, isIP bool) (bool, error)
	DeleteDevice(ip_address string) error
}

type User interface {
	GetUserByID(id int) (app.User, error)
	GetRole(username string) (string, error)
}

type Info interface {
	PingDevices() ([][]string, error)
	SaveMinerData(data app.MinerData, ip_address string) error
	Transform(devices []app.MinerDevice) (map[string][]app.MinerData, error)
	GetCharacteristicsHistory(device_ip string) ([]app.MinerData, error)
	DetermineIP(mac_address string) string
	DetermineMAC(ip_address string) string
	SaveAvailableAddresses(list [][]string) error
}

type Comment interface {
	Comment(ip_address, username, comment string) error
	GetCommentsHistory(ip_address string) ([]app.Comment, error)
	DeleteComment(creation_date time.Time) error
	EditComment(creation_date time.Time, newContent string) error
}

type Chat interface {
	SaveMessage(message app.Message) error
	ReadUserMessages(sender string) ([]app.Message, error)
	//ReadOperatorMessages(recipient string) ([]app.Message, error)
	GetSenders() ([]string, error)
}

type Service struct {
	Authorization
	Miner
	User
	Info
	Comment
	Chat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Miner:         NewMinerService(repos.Miner),
		Info:          NewInfoService(repos.Info),
		Comment:       NewCommentService(repos.Comment),
		Chat:          NewChatService(repos.Chat),
	}
}
