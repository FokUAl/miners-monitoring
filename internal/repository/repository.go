package repository

import (
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
}

type Miner interface {
	GetDevice(address string) (app.MinerDevice, error)
	GetAllDevices() ([]app.MinerDevice, error)
	GetDevicesInfo() ([]app.AddInfo, error)
	AddNew(dev app.MinerDevice) error
	IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error)
	GetDevicesByType(miner_type string) ([]app.MinerDevice, error)
	GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error)
	GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error)
	GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error)
	GetDevicesByUser(username string) ([]app.MinerDevice, error)
	UpdateDevice(newInfo app.AddInfo) error
	IsAddressFree(address string, isIP bool) (bool, error)
	DeleteDevice(ip_address string) error
}

type User interface {
	GetUserByID(id int) (app.User, error)
	GetUserRole(username string) (string, error)
}

type Info interface {
	SaveMinerData(data app.MinerData, ip_address string) error
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
	//ReadUserMessages(sender string) ([]app.Message, error)
	//ReadOperatorMessages(recipient string) ([]app.Message, error)
	GetSenders() ([]string, error)
}

type Repository struct {
	Authorization
	User
	Miner
	Info
	Comment
	Chat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Miner:         NewMinerPostgres(db),
		Info:          NewInfoPostgres(db),
		Comment:       NewCommentPostgres(db),
		Chat:          NewChatPostgres(db),
	}
}
