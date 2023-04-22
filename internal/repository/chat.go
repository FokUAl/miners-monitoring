package repository

import (
	"fmt"

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
	statement := `INSERT INTO chat_history (creation_date, content, sender, sender_role, 
		recipient) VALUES ($1, $2, $3, $4, $5)`

	_, err := p.db.Exec(statement, message.Time, message.Content,
		message.Sender, message.SenderRole, message.Recipient)

	return err
}

// return a list of users that send messages to operators
func (p *ChatPostgres) GetSenders() ([]string, error) {
	statement := `SELECT sender FROM chat_history WHERE sender_role = $1`

	rows, err := p.db.Query(statement, "User")
	if err != nil {
		return nil, fmt.Errorf("GetSenders: %w", err)
	}

	defer rows.Close()

	var result []string
	for rows.Next() {
		var sender string
		err = rows.Scan(&sender)
		if err != nil {
			return nil, fmt.Errorf("GetSenders: %w", err)
		}

		result = append(result, sender)
	}

	return result, nil
}

func (p *ChatPostgres) ReadUserMessages(sender string) ([]app.Message, error) {
	query := `SELECT creation_date, content, sender, recipient, is_read
	FROM chat_history WHERE sender = $1`

	rows, err := p.db.Query(query, sender)
	if err != nil {
		return nil, fmt.Errorf("ReadUserMessages: %w", err)
	}

	defer rows.Close()

	var result []app.Message
	for rows.Next() {
		var message app.Message
		err = rows.Scan(&message.Time, &message.Content,
			&message.Sender, &message.Recipient, &message.IsRead)
		if err != nil {
			return nil, fmt.Errorf("ReadUserMessages: %w", err)
		}

		result = append(result, message)
	}

	return result, nil
}
