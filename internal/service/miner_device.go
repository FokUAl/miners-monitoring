package service

import (
	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type MinerService struct {
	repo repository.Miner
}

func NewMinerService(repo repository.Miner) *MinerService {
	return &MinerService{
		repo: repo,
	}
}

func (s *MinerService) GetDevice(address string) (app.MinerDevice, error) {
	return s.repo.GetDevice(address)
}

func (s *MinerService) GetAllDevices() ([]app.MinerDevice, error) {
	return s.repo.GetAllDevices()
}

func (s *MinerService) AddNew(dev app.MinerDevice) error {
	return s.repo.AddNew(dev)
}

// Mapping device IP with location and owner
func (s *MinerService) MappDevices(mappingInfo []app.AddInfo) error {
	for i := 0; i < len(mappingInfo); i++ {
		err := s.repo.UpdateDevice(mappingInfo[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *MinerService) IsDeviceAdded(address string) (bool, error) {
	device, err := s.repo.GetDevice(address)

	return device != app.MinerDevice{}, err
}

func (s *MinerService) IsLocationFree(shelfNum, rowNum, columnNum int) (bool, error) {
	return s.repo.IsLocationFree(shelfNum, rowNum, columnNum)
}

func (s *MinerService) GetDevicesByType(miner_type string) ([]app.MinerDevice, error) {
	return s.repo.GetDevicesByType(miner_type)
}

func (s *MinerService) GetDevicesByStatus(miner_status string) ([]app.MinerDevice, error) {
	return s.repo.GetDevicesByStatus(miner_status)
}

func (s *MinerService) GetDevicesByCoin(coin_type string) ([]app.MinerDevice, error) {
	return s.repo.GetDevicesByCoin(coin_type)
}

func (s *MinerService) GetDeviceByLocation(shelf int, column int, row int) (app.MinerDevice, error) {
	return s.repo.GetDeviceByLocation(shelf, column, row)
}

func (s *MinerService) GetDevicesByUser(username string) ([]app.MinerDevice, error) {
	return s.repo.GetDevicesByUser(username)
}
