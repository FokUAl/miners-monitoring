package service

import (
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

func (s *MinerService) AddDevices(model string, isIP bool, connections, shelfData, rowData, columnData []string) error {
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

		shelfInt, err := strconv.Atoi(shelfData[i])
		if err != nil {
			return err
		}
		device.Shelf = shelfInt

		rowInt, err := strconv.Atoi(rowData[i])
		if err != nil {
			return err
		}
		device.Row = rowInt

		colInt, err := strconv.Atoi(columnData[i])
		if err != nil {
			return err
		}
		device.Column = colInt

		err = s.AddNew(device)
		if err != nil {
			return err
		}
	}

	return nil
}
