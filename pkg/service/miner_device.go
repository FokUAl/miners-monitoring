package service

import (
	"fmt"
	"strconv"

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

func (s *MinerService) GetDevice(ip_address string) (app.MinerDevice, error) {
	return s.repo.GetDevice(ip_address)
}

func (s *MinerService) GetAllDevices() ([]app.MinerDevice, error) {
	return s.repo.GetAllDevices()
}

func (s *MinerService) AddNew(dev app.MinerDevice) error {
	return s.repo.AddNew(dev)
}

func (s *MinerService) AddDevices(model string, isIP bool, connections []string, locInfo [][]string) error {
	for i := 0; i < len(connections); i++ {
		// check to existence of device
		existedDevice, err := s.GetDevice(connections[i])
		if isIP && err != nil {
			return fmt.Errorf("AddDevices: device is not exist")
		}

		shelfNum, err := strconv.Atoi(locInfo[0][i])
		if err != nil {
			return err
		}
		rowNum, err := strconv.Atoi(locInfo[1][i])
		if err != nil {
			return err
		}
		columnNum, err := strconv.Atoi(locInfo[2][i])
		if err != nil {
			return err
		}
		// change device location
		existedDevice.Shelf = shelfNum
		existedDevice.Row = rowNum
		existedDevice.Column = columnNum

		err = s.AddNew(existedDevice)
		if err != nil {
			return err
		}

	}

	return nil
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
