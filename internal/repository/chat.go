package repository

import (
	"fmt"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/pkg"
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

		if !pkg.IsContain(sender, result) {
			result = append(result, sender)
		}
	}

	return result, nil
}

func (p *ChatPostgres) ReadUserMessages(sender string) ([]app.Message, error) {
	query := `SELECT creation_date, content, sender, recipient, 
	is_read, sender_role FROM chat_history 
	WHERE sender = $1 OR recipient = $2
	ORDER BY creation_date ASC`

	rows, err := p.db.Query(query, sender, sender)
	if err != nil {
		return nil, fmt.Errorf("ReadUserMessages: %w", err)
	}

	defer rows.Close()

	var result []app.Message
	for rows.Next() {
		var message app.Message
		err = rows.Scan(&message.Time, &message.Content,
			&message.Sender, &message.Recipient, &message.IsRead, &message.SenderRole)
		if err != nil {
			return nil, fmt.Errorf("ReadUserMessages: %w", err)
		}

		result = append(result, message)
	}

	return result, nil
}

func (p *ChatPostgres) UpdateMessageStatus(sender, recipient string) error {
	query := `UPDATE chat_history SET recipient = $1, is_read = TRUE 
		WHERE sender = $2 AND sender_role = 'User'`

	_, err := p.db.Exec(query, recipient, sender)

	return err
}

func (p *ChatPostgres) ReadMessages(sender, recipient string) ([]app.Message, error) {
	err := p.UpdateMessageStatus(sender, recipient)
	if err != nil {
		return nil, fmt.Errorf("UpdateMessageStatus: %w", err)
	}

	query := `SELECT creation_date, content, sender, recipient, 
	is_read, sender_role FROM chat_history 
	WHERE (sender = $1 AND recipient = $2) OR (sender = $3 AND recipient = $4)
	ORDER BY creation_date ASC`

	rows, err := p.db.Query(query, sender, recipient, recipient, sender)
	if err != nil {
		return nil, err
	}

	var result []app.Message
	for rows.Next() {
		var message app.Message
		err = rows.Scan(&message.Time, &message.Content,
			&message.Sender, &message.Recipient, &message.IsRead, &message.SenderRole)
		if err != nil {
			return nil, fmt.Errorf("ReadMessages: %w", err)
		}

		result = append(result, message)
	}

	return result, nil

}
