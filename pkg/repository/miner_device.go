package repository

import (
	"fmt"

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

func (p *MinerPostgres) GetDevice(id int) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, area, miner_status, coin,
		ip_address, mac_address FROM miner_devices WHERE id = $1`

	err := p.db.Get(&device, query, id)

	return device, err
}

func (p *MinerPostgres) GetAllDevices() ([]app.MinerDevice, error) {
	var devices []app.MinerDevice

	query := `SELECT miner_type, shelf, _row, column, miner_status, coin, 
		ip_address, mac_address FROM miner_devices`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllDevices: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var device app.MinerDevice

		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.MinerStatus,
			&device.Coin, &device.IPAddress, &device.MACAddress)
		if err != nil {
			return nil, fmt.Errorf("GetAllDevices: %w", err)
		}

		devices = append(devices, device)
	}

	return devices, nil
}
