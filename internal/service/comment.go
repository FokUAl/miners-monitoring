package service

import (
	"time"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Comment(ip_address, username, comment string) error {
	return s.repo.Comment(ip_address, username, comment)
}

func (s *CommentService) GetCommentsHistory(ip_address string) ([]app.Comment, error) {
	return s.repo.GetCommentsHistory(ip_address)
}

func (s *CommentService) DeleteComment(creation_date time.Time) error {
	return s.repo.DeleteComment(creation_date)
}

func (s *CommentService) EditComment(creation_date time.Time, newContent string) error {
	return s.repo.EditComment(creation_date, newContent)
}
