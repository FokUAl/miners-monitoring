package service

import (
	"fmt"

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

func (s *MinerService) AddDevices(model string, isIP bool, connections []string) error {
	for i := 0; i < len(connections); i++ {
		var device app.MinerDevice
		device.MinerType = model

		if isIP {
			device.IPAddress = connections[i]
			device.MACAddress = "-"
		} else {
			device.MACAddress = connections[i]
			device.IPAddress = "-"
		}

		// check to existence of device
		existedDevice, err := s.GetDevice(connections[i])
		if isIP && err != nil {
			return fmt.Errorf("AddDevices: device is not exist: %w", err)
		}

		err = s.AddNew(existedDevice)
		if err != nil {
			return err
		}

	}

	return nil
}
