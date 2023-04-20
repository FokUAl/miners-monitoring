package repository

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{
		db: db,
	}
}

func (p *ChatPostgres) SaveMessage(message app.Message) error {
	statement := `INSERT INTO chat_history (creation_date, content, sender, recipient) 
	VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(statement, message.Time, message.Content,
		message.Sender, message.Recipient)

	return err
}
