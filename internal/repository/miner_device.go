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

func (p *MinerPostgres) GetDevicesInfo() ([]app.AddInfo, error) {
	var devicesInfo []app.AddInfo

	query := `SELECT miner_type, ip_address, shelf, _row, col, owner_ FROM miner_devices`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesInfo: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var info app.AddInfo
		err = rows.Scan(&info.MinerType, &info.IP, &info.Shelf, &info.Row, &info.Column, &info.Owner)
		if err != nil {
			return nil, fmt.Errorf("GetDevicesInfo: %w", err)
		}

		devicesInfo = append(devicesInfo, info)
	}

	return devicesInfo, nil
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

		query = `SELECT elapsed, mhs_av, temperature, fan_speed_in, 
			fan_speed_out, power_mode, chip_temp_min, chip_temp_max, chip_temp_avg
			FROM miner_characteristics WHERE ip_address = $1`

		err = p.db.QueryRow(query, device.IPAddress).Scan(&device.Characteristics.Elapsed,
			&device.Characteristics.MHSav, &device.Characteristics.Temperature,
			&device.Characteristics.FanSpeedIn, &device.Characteristics.FanSpeedOut,
			&device.Characteristics.PowerMode, &device.Characteristics.ChipTempMin,
			&device.Characteristics.ChipTempMax, &device.Characteristics.ChipTempAvg)

		if err != nil {
			return nil, fmt.Errorf("GetAllDevices: %w", err)
		}

		device.Characteristics.IP = device.IPAddress

		devices = append(devices, device)
	}

	return devices, nil
}

func (p *MinerPostgres) AddNew(dev app.MinerDevice) error {
	query := `INSERT INTO miner_devices (shelf, _row, col, owner_, ip_address, miner_type) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(query, dev.Shelf, dev.Row, dev.Column, dev.Owner, dev.IPAddress, dev.MinerType)

	if err != nil {
		return err
	}

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
		ip_address, mac_address FROM miner_devices 
		WHERE shelf = $1 AND _row = $2 AND col = $3`

	err := p.db.QueryRow(query, shelfNum, rowNum, columnNum).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.Characteristics.MAC)

	return device == app.MinerDevice{}, err
}

func (p *MinerPostgres) IsIPFree(ip_address string) (bool, error) {
	var device app.MinerDevice

	query := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin,
		ip_address, mac_address FROM miner_devices 
		WHERE ip_address = $1`

	err := p.db.QueryRow(query, ip_address).Scan(&device.MinerType, &device.Shelf, &device.Row,
		&device.Column, &device.Owner, &device.MinerStatus, &device.Coin, &device.IPAddress,
		&device.Characteristics.MAC)

	return device == app.MinerDevice{}, err
}

func (p *MinerPostgres) GetDevicesByType(miner_type string) ([]app.MinerDevice, error) {
	var result []app.MinerDevice

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address
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

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address
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

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address
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
	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address
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

	statement := `SELECT miner_type, shelf, _row, col, owner_, miner_status, coin, ip_address, mac_address
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
	query := `UPDATE miner_devices SET owner_ = $1, shelf = $2, _row = $3, col = $4,
		miner_type = $5 WHERE ip_address = $6`

	_, err := p.db.Exec(query, newInfo.Owner, newInfo.Shelf, newInfo.Row,
		newInfo.Column, newInfo.MinerType, newInfo.IP)
	return err
}

func (p *MinerPostgres) DeleteDevice(ip_address string) error {
	query := `DELETE FROM miner_devices WHERE ip_address = $1`

	_, err := p.db.Exec(query, ip_address)
	return err
}
