package repository

import (
	"fmt"
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/jmoiron/sqlx"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{
		db: db,
	}
}

func (p *CommentPostgres) Comment(ip_address, username, comment string) error {
	query := `INSERT INTO comments (ip_address, username, comment, creation_date)
		VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(query, ip_address, username, comment, time.Now())

	return err
}

func (p *CommentPostgres) GetCommentsHistory(ip_address string) ([]app.Comment, error) {
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

func (p *CommentPostgres) DeleteComment(creation_date time.Time) error {
	query := `DELETE FROM comments WHERE creation_date = $1`

	_, err := p.db.Exec(query, creation_date)
	return err
}

func (p *CommentPostgres) EditComment(creation_date time.Time, newContent string) error {
	query := `UPDATE comments SET creation_date = $1, comment = $2
		WHERE creation_date = $3`
	_, err := p.db.Exec(query, time.Now(), newContent, creation_date)
	return err
}
