package repository

import (
	"database/sql"
	"fmt"
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{
		db: db,
	}
}

func (p *InfoPostgres) SaveMinerData(data app.MinerData, ip_address string) error {
	query := `INSERT INTO miner_characteristics (elapsed, mhs_av, temperature, fan_speed1, 
		fan_speed2, fan_speed3, fan_speed4, power_mode, chip_temp_min, chip_temp_max, chip_temp_avg, 
		creation_date, ip_address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := p.db.Exec(query, data.Elapsed, data.MHSav, data.Temperature, data.FanSpeed1,
		data.FanSpeed2, data.FanSpeed3, data.FanSpeed4, data.PowerMode, data.ChipTempMin, data.ChipTempMax,
		data.ChipTempAvg, time.Now(), ip_address)
	return err
}

func (p *InfoPostgres) GetCharacteristicsHistory(device_ip string) ([]app.MinerData, error) {
	query := `SELECT elapsed, mhs_av, temperature, fan_speed1, fan_speed2, 
		fan_speed3, fan_speed4, power_mode, chip_temp_min,
		chip_temp_max, chip_temp_avg, creation_date FROM miner_characteristics
		WHERE ip_address = $1 ORDER BY creation_date ASC`

	rows, err := p.db.Query(query, device_ip)
	if err != nil {
		return nil, fmt.Errorf("GetCharacteristicsHistory: %w", err)
	}

	defer rows.Close()

	var result []app.MinerData
	for rows.Next() {
		var record app.MinerData

		err = rows.Scan(&record.Elapsed, &record.MHSav, &record.Temperature, &record.FanSpeed1, &record.FanSpeed2,
			&record.FanSpeed3, &record.FanSpeed4, &record.PowerMode, &record.ChipTempMin,
			&record.ChipTempMax, &record.ChipTempAvg, &record.Time)
		if err != nil {
			return nil, fmt.Errorf("GetCharacteristicsHistory: %w", err)
		}

		result = append(result, record)
	}

	return result, nil
}

func (p *InfoPostgres) DetermineIP(mac_address string) string {
	var result string
	statement := `SELECT ip_address from mac_ip WHERE mac_address = $1`

	err := p.db.QueryRow(statement, mac_address).Scan(&result)
	if err == sql.ErrNoRows {
		return ""
	}

	return result
}

func (p *InfoPostgres) DetermineMAC(ip_address string) string {
	var result string
	statement := `SELECT mac_address from mac_ip WHERE ip_address = $1`

	err := p.db.QueryRow(statement, ip_address).Scan(&result)
	if err == sql.ErrNoRows {
		return ""
	}

	return result
}

func (p *InfoPostgres) SaveAvailableAddresses(list [][]string) error {
	query := `DELETE FROM mac_ip`

	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	statement := `INSERT INTO mac_ip (ip_address, mac_address)
		VALUES ($1, $2)`

	for _, elem := range list {
		_, err := p.db.Exec(statement, elem[1], elem[0])
		if err != nil {
			return err
		}
	}

	return nil
}
