package service

import app "github.com/FokUAl/miners-monitoring"

func (s *AuthService) GetUser(username, password string) (app.User, error) {
	return s.repo.GetUser(username, password)
}
