package repository

import (
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
