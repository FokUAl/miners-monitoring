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

func (p *MinerPostgres) GetDevice(address string, isIP bool) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM all_devices WHERE ip_address = $1`

	if !isIP {
		query = `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM all_devices WHERE mac_address = $1`
	}

	err := p.db.QueryRow(query, address).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.MACAddress, &device.Pool)

	return device, err
}

func (p *MinerPostgres) GetAllDevices() ([]app.MinerDevice, error) {
	var devices []app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, 
		ip_address, mac_address FROM miner_devices`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllDevices: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var device app.MinerDevice

		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.Owner,
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.MACAddress)
		if err != nil {
			return nil, fmt.Errorf("GetAllDevices: %w", err)
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (p *MinerPostgres) AddNew(dev app.MinerDevice) error {
	query := `INSERT INTO miner_devices (miner_type, shelf, _row, col, owner_, miner_status,
		coin, ip_address, mac_address, _pool) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.Exec(query, dev.MinerType, dev.Shelf, dev.Row, dev.Column, dev.Owner, dev.MinerStatus,
		dev.Coin, dev.IPAddress, dev.MACAddress, dev.Pool)

	return err
}

func (p *MinerPostgres) GetDeviceFromDB(address string, isIP bool) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM miner_devices WHERE ip_address = $1`
	if !isIP {
		query = `SELECT miner_type, shelf, _row, col, miner_status, coin,
		ip_address, mac_address, _pool FROM miner_devices WHERE mac_address = $1`
	}

	err := p.db.QueryRow(query, address).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.MACAddress, &device.Pool)

	return device, err
}

func (p *MinerPostgres) IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM miner_devices 
		WHERE shelf = $1 AND _row = $2 AND col = $3`

	err := p.db.QueryRow(query, shelfNum, rowNum, columnNum).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.MACAddress, &device.Pool)

	return device == app.MinerDevice{}, err
}

func (p *MinerPostgres) GetDevicesByType(miner_type string) ([]app.MinerDevice, error) {
	var result []app.MinerDevice

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address, _pool 
		FROM miner_devices WHERE miner_type = $1`
	rows, err := p.db.Query(statement, miner_type)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesByType: %w", err)
	}

	for rows.Next() {
		var device app.MinerDevice
		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.Owner,
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.MACAddress, &device.Pool)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesByType: %w", err)
		}

		result = append(result, device)
	}

	return result, nil
}

func (p *MinerPostgres) GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error) {
	var result []app.MinerDevice

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address, _pool 
		FROM miner_devices WHERE coin = $1`
	rows, err := p.db.Query(statement, coin_type)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesByCoin: %w", err)
	}

	for rows.Next() {
		var device app.MinerDevice
		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.Owner,
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.MACAddress, &device.Pool)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesByCoin: %w", err)
		}

		result = append(result, device)
	}

	return result, nil
}

func (p *MinerPostgres) GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error) {
	var result []app.MinerDevice

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address, _pool 
		FROM miner_devices WHERE miner_status = $1`
	rows, err := p.db.Query(statement, miner_status)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesByStatus: %w", err)
	}

	for rows.Next() {
		var device app.MinerDevice
		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.Owner,
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.MACAddress, &device.Pool)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesByStatus: %w", err)
		}

		result = append(result, device)
	}

	return result, nil
}

func (p *MinerPostgres) GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error) {
	var result app.MinerDevice
	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address, _pool 
		FROM miner_devices WHERE shelf = $1 and _row = $2 and col = $3`
	err := p.db.QueryRow(statement, shelf, row, column).Scan(&result.MinerType, &result.Shelf, &result.Row,
		&result.Column, &result.Owner, &result.MinerStatus, &result.Coin, &result.IPAddress, &result.MACAddress, &result.Pool)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *MinerPostgres) GetDevicesByUser(username string) ([]app.MinerDevice, error) {
	var result []app.MinerDevice

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address, _pool 
		FROM miner_devices WHERE owner_ = $1`

	rows, err := p.db.Query(statement, username)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesByStatus: %w", err)
	}

	for rows.Next() {
		var device app.MinerDevice
		err = rows.Scan(&device.MinerType, &device.Shelf, &device.Row, &device.Column, &device.Owner,
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.MACAddress, &device.Pool)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesByUser: %w", err)
		}

		result = append(result, device)
	}

	return result, nil
}
