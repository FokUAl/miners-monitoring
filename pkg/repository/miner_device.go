package repository

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type MinerPostgres struct {
	db *sqlx.DB
}

func NewMinerPostgres(db *sqlx.DB) *MinerPostgres {
	return &MinerPostgres{
		db: db,
	}
}

func (r *MinerPostgres) GetDevice(id int) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, area, miner_status, coin,
		ip_address, mac_address FROM miner_devices WHERE id = $1`

	err := r.db.Get(&device, query, id)

	return device, err
}
