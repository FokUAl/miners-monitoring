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
	GetDevice(id int) (app.MinerDevice, error)
}

type User interface{}

type Repository struct {
	Authorization
	User
	Miner
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Miner:         NewMinerPostgres(db),
	}
}
