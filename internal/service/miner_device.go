package service

import (
	"database/sql"
	"fmt"
	"strconv"

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

func (s *MinerService) GetDevice(address string, isIP bool) (app.MinerDevice, error) {
	return s.repo.GetDevice(address, isIP)
}

func (s *MinerService) GetAllDevices() ([]app.MinerDevice, error) {
	return s.repo.GetAllDevices()
}

func (s *MinerService) AddNew(dev app.MinerDevice) error {
	return s.repo.AddNew(dev)
}

func (s *MinerService) AddDevices(model string, isIP bool, connections []string, locInfo [][]string) error {
	for i := 0; i < len(connections); i++ {
		// check to existence of device physically
		existedDevice, err := s.repo.GetDevice(connections[i], isIP)
		if err != nil {
			return fmt.Errorf("device is not exist: %s", connections[i])
		}

		// check device to existance in database
		isAdded, err := s.IsDeviceAdded(connections[i], isIP)
		if err != sql.ErrNoRows && err != nil {
			return err
		} else if isAdded {
			return fmt.Errorf("device has already been added: %s", connections[i])
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

		// is location free
		isFree, err := s.IsLocationFree(shelfNum, rowNum, columnNum)
		if err != sql.ErrNoRows && err != nil {
			return err
		} else if !isFree {
			return fmt.Errorf("location isn't free: %d-%d-%d",
				shelfNum, columnNum, rowNum)
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

func (s *MinerService) IsDeviceAdded(address string, isIP bool) (bool, error) {
	device, err := s.repo.GetDeviceFromDB(address, isIP)

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
