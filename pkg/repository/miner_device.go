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

func (p *MinerPostgres) GetDevice(ip_address string) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, miner_status, coin,
		ip_address, mac_address, _pool FROM all_devices WHERE ip_address = $1`

	err := p.db.QueryRow(query, ip_address).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.MACAddress, &device.Pool)

	return device, err
}

func (p *MinerPostgres) GetAllDevices() ([]app.MinerDevice, error) {
	var devices []app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, miner_status, coin, 
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

func (p *MinerPostgres) AddNew(dev app.MinerDevice) error {
	query := `INSERT INTO miner_devices (miner_type, shelf, _row, col, miner_status,
		coin, ip_address, mac_address, _pool) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.Exec(query, dev.MinerType, dev.Shelf, dev.Row, dev.Column, dev.MinerStatus,
		dev.Coin, dev.IPAddress, dev.MACAddress, dev.Pool)

	return err
}
