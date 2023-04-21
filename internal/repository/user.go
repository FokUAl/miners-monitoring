package repository

import (
	"fmt"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (p *UserPostgres) GetUserByID(id int) (app.User, error) {
	var user app.User
	query := `SELECT id, username, role_ FROM users WHERE id = $1`
	err := p.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Role)

	return user, err
}

func (p *UserPostgres) GetUserRole(username string) (string, error) {
	var result string
	query := `SELECT role_ FROM users WHERE username = $1`
	err := p.db.QueryRow(query, username).Scan(&result)

	return result, err
}

func (p *UserPostgres) ListOfUsers() ([]string, error) {
	var result []string
	query := `SELECT username FROM users WHERE role_ = $1`

	rows, err := p.db.Query(query, "User")
	if err != nil {
		return nil, fmt.Errorf("ListOfUsers: %w", err)
	}

	for rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			return nil, fmt.Errorf("ListOfUsers: %w", err)
		}

		result = append(result, username)
	}

	return result, nil

}
