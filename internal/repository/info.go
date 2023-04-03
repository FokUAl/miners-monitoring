package repository

import (
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
	query := `INSERT INTO miner_characteristics (elapsed, mhs_av, temperature, fan_speed_in, 
		fan_speed_out, power_mode, chip_temp_min, chip_temp_max, chip_temp_avg, 
		creation_date, ip_address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := p.db.Exec(query, data.Elapsed, data.MHSav, data.Temperature, data.FanSpeedIn,
		data.FanSpeedOut, data.PowerMode, data.ChipTempMin, data.ChipTempMax,
		data.ChipTempAvg, time.Now(), ip_address)
	return err
}

func (p *InfoPostgres) GetCharacteristicsHistory(device_ip string) ([]app.MinerData, error) {
	query := `SELECT elapsed, mhs_av, temperature, fan_speed_in, fan_speed_out, power_mode, chip_temp_min,
		chip_temp_max, chip_temp_avg, creation_date FROM miner_characteristics
		WHERE ip_address = $1`

	rows, err := p.db.Query(query, device_ip)
	if err != nil {
		return nil, fmt.Errorf("GetCharacteristicsHistory: %w", err)
	}

	defer rows.Close()

	var result []app.MinerData
	for rows.Next() {
		var record app.MinerData

		err = rows.Scan(&record.Elapsed, &record.MHSav, &record.Temperature, &record.FanSpeedIn, &record.FanSpeedOut,
			&record.PowerMode, &record.ChipTempMin, &record.ChipTempMax, &record.ChipTempAvg, &record.Time)
		if err != nil {
			return nil, fmt.Errorf("GetCharacteristicsHistory: %w", err)
		}

		result = append(result, record)
	}

	return result, nil
}
