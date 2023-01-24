package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/pkg/repository"
)

type MinerService struct {
	repo repository.Miner
}

func NewMinerService(repo repository.Miner) *MinerService {
	return &MinerService{
		repo: repo,
	}
}

func (s *MinerService) GetDevice(id int) (app.MinerDevice, error) {
	return s.repo.GetDevice(id)
}

func (s *MinerService) GetAllDevices() ([]app.MinerDevice, error) {
	return s.repo.GetAllDevices()
}
