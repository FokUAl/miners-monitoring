package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) SaveMessage(message app.Message) error {
	return s.repo.SaveMessage(message)
}

func (s *ChatService) GetSenders() ([]string, error) {
	return s.repo.GetSenders()
}

func (s *ChatService) ReadUserMessages(sender string) ([]app.Message, error) {
	return s.repo.ReadUserMessages(sender)
}

func (s *ChatService) ReadMessages(sender, recipient string) ([]app.Message, error) {
	return s.repo.ReadMessages(sender, recipient)
}
