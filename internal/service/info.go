package service

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{
		repo: repo,
	}
}

func (s *InfoService) ParsingData(data string) (app.MinerData, error) {
	var minerData app.MinerData

	r, err := regexp.Compile("'[A-Za-z0-9% ]+': ('?[0-9A-Za-z:._ -]+'?)")
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't compile regexp: %s", err.Error())
	}

	arr := r.FindAllString(data, -1)
	data_map := make(map[string]string)
	for _, val := range arr {
		keyvalue := strings.Split(val, ": ")
		key := strings.Trim(keyvalue[0], "'")
		value := strings.Trim(keyvalue[1], "'")
		data_map[key] = value
	}

	minerData.MAC = data_map["MAC"]
	minerData.PowerMode = data_map["Power Mode"]
	minerData.Temperature, err = strconv.ParseFloat(data_map["Temperature"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temperature: %s", err.Error())
	}

	minerData.MHSav, err = strconv.ParseFloat(data_map["MHS av"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse MHS av: %s", err.Error())
	}

	minerData.FanSpeedIn, err = strconv.ParseInt(data_map["Fan Speed In"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse MHS av: %s", err.Error())
	}

	minerData.FanSpeedOut, err = strconv.ParseInt(data_map["Fan Speed Out"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed Out: %s", err.Error())
	}

	minerData.Elapsed, err = strconv.ParseInt(data_map["Elapsed"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Elapsed: %s", err.Error())
	}

	minerData.ChipTempAvg, err = strconv.ParseFloat(data_map["Chip Temp Avg"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Chip Temp Avg: %s", err.Error())
	}

	minerData.ChipTempMax, err = strconv.ParseFloat(data_map["Chip Temp Max"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Chip Temp Max: %s", err.Error())
	}

	minerData.ChipTempMin, err = strconv.ParseFloat(data_map["Chip Temp Min"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Chip Temp Min: %s", err.Error())
	}

	return minerData, nil
}

func (s *InfoService) PingDevices() ([]string, error) {
	cmd := exec.Command("arp", "-an")

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("pingDevices: %s", err.Error())
	}

	r, err := regexp.Compile(`\(.+\)`)
	if err != nil {
		return nil, fmt.Errorf("pingDevices: can't compile regexp: %s", err.Error())
	}

	result := r.FindAllString(string(out), -1)
	for ind := 0; ind < len(result); ind++ {
		result[ind] = strings.Trim(result[ind], "()")
	}

	return result, nil
}

// func checks is it actually asic by IP
func (s *InfoService) CheckResponse(response string) error {
	// search target segment of string
	regEx, err := regexp.Compile(`'STATUS': '[a-zA-Z]+'`)
	if err != nil {
		return err
	}
	errText := regEx.FindAllString(response, -1)
	if len(errText) != 1 {
		return fmt.Errorf("unknown response %s", response)
	} else if errText[0] == "'STATUS': 'error'" {
		return fmt.Errorf("error response %s", response)
	}

	return nil
}

func (s *InfoService) SaveMinerData(data app.MinerData) error {
	return s.repo.SaveMinerData(data)
}
