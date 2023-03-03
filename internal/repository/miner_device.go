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
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.Characteristics.MAC)
		if err != nil {
			return nil, fmt.Errorf("GetAllDevices: %w", err)
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (p *MinerPostgres) AddNew(dev app.MinerDevice) error {
	query := `INSERT INTO miner_devices (miner_type, shelf, _row, col, owner_, miner_status,
		coin, ip_address, mac_address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.Exec(query, dev.MinerType, dev.Shelf, dev.Row, dev.Column, dev.Owner, dev.MinerStatus,
		dev.Coin, dev.IPAddress, dev.Characteristics.MAC)

	if err != nil {
		return err
	}

	query = `INSERT INTO miner_characteristics (elapsed, mhs_av, temperature, fan_speed_in, fan_speed_out, power_mode,
		chip_temp_min, chip_temp_max, chip_temp_avg) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = p.db.Exec(query, dev.Characteristics.Elapsed, dev.Characteristics.MHSav,
		dev.Characteristics.Temperature, dev.Characteristics.FanSpeedIn, dev.Characteristics.FanSpeedOut,
		dev.Characteristics.PowerMode, dev.Characteristics.ChipTempMin, dev.Characteristics.ChipTempMax,
		dev.Characteristics.ChipTempAvg)

	return err
}

func (p *MinerPostgres) GetDevice(address string) (app.MinerDevice, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM miner_devices WHERE ip_address = $1`

	err := p.db.QueryRow(query, address).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.Characteristics.MAC)

	return device, err
}

func (p *MinerPostgres) IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address, _pool FROM miner_devices 
		WHERE shelf = $1 AND _row = $2 AND col = $3`

	err := p.db.QueryRow(query, shelfNum, rowNum, columnNum).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.Characteristics.MAC)

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
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.Characteristics.MAC)
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
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.Characteristics.MAC)
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
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.Characteristics.MAC)
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
		&result.Column, &result.Owner, &result.MinerStatus, &result.Coin, &result.IPAddress, &result.Characteristics.MAC)
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
			&device.MinerStatus, &device.Coin, &device.IPAddress, &device.Characteristics.MAC)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesByUser: %w", err)
		}

		result = append(result, device)
	}

	return result, nil
}

func (p *MinerPostgres) UpdateDevice(newInfo app.AddInfo) error {
	query := `UPDATE miner_devices SET owner_ = $1, shelf = $2, _row = $3, col = $4
		WHERE ip_address = $5`

	_, err := p.db.Exec(query, newInfo.Owner, newInfo.Shelf, newInfo.Row,
		newInfo.Column, newInfo.IP)
	return err
}
