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

// TO DO
// Add user role
func (p *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := `INSERT INTO users (email, role_, username, password_hash) VALUES ($1, 'User', $2, $3)
		RETURNING id`

	row := p.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *AuthPostgres) GetUser(username, password string) (app.User, error) {
	var user app.User
	query := `SELECT id, role_ FROM users WHERE username = $1 AND password_hash = $2`
	err := p.db.QueryRow(query, username, password).Scan(&user.Id, &user.Role)

	return user, err
}
