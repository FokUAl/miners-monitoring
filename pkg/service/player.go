package service

import app "github.com/HgCl2/rock_scissors_paper"

func (s *AuthService) GetUser(username, password string) (app.Player, error) {
	return s.repo.GetUser(username, password)
}
