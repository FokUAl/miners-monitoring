package repository

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
}

type Miner interface {
	GetDevice(address string, isIP bool) (app.MinerDevice, error)
	GetAllDevices() ([]app.MinerDevice, error)
	AddNew(dev app.MinerDevice) error
	GetDeviceFromDB(address string, isIP bool) (app.MinerDevice, error)
	IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error)
	GetDevicesByType(miner_type string) ([]app.MinerDevice, error)
	GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error)
	GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error)
	GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error)
	GetDevicesByUser(username string) ([]app.MinerDevice, error)
}

type User interface {
	GetUserByID(id int) (app.User, error)
}

type Repository struct {
	Authorization
	User
	Miner
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Miner:         NewMinerPostgres(db),
	}
}
