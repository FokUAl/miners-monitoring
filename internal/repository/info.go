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

func (p *InfoPostgres) Comment(ip_address, username, comment string) error {
	query := `INSERT INTO comments (ip_address, username, comment, creation_date)
		VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(query, ip_address, username, comment, time.Now())

	return err
}

func (p *InfoPostgres) GetCommentsHistory(ip_address string) ([]app.Comment, error) {
	query := `SELECT creation_date, username, comment FROM comments
		WHERE ip_address = $1`

	rows, err := p.db.Query(query, ip_address)
	if err != nil {
		return nil, fmt.Errorf("GetCommentsHistory: %w", err)
	}

	defer rows.Close()

	var result []app.Comment
	for rows.Next() {
		var tempComment app.Comment

		err = rows.Scan(&tempComment.CreationDate, &tempComment.Username, &tempComment.Content)
		if err != nil {
			return nil, fmt.Errorf("GetCommentsHistory: %w", err)
		}

		result = append(result, tempComment)
	}

	return result, nil
}
