package repository

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := `INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)
		RETURNING id`

	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (app.User, error) {
	var user app.User
	query := `SELECT id FROM users WHERE username = $1 AND password_hash = $2`
	err := r.db.Get(&user, query, username, password)

	return user, err
}
