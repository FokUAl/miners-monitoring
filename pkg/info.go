package pkg

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/kluctl/go-embed-python/python"
)

// Uses python script to send request to cgminer interface
// and return response from it as string.
func GetAsicInfo(ip string, command string) (string, error) {
	ep, err := python.NewEmbeddedPython("example")
	if err != nil {
		return "", err
	}

	cmd := ep.PythonCmd("./pkg/cgminer.py", command, ip)

	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = w

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// Func for getting data by IP and
// transforming data to MinerData struct
func ResponseToStruct(ip_address string) (app.MinerData, error) {
	var result app.MinerData

	response, err := GetAsicInfo(ip_address, "summary")
	if err != nil {
		return result, err
	}

	err = CheckResponse(response)
	if err != nil {
		return result, err
	}

	result, err = ParsingData(response)

	return result, err
}

// Func for getting data by IP and
// transforming data to MinerData struct
// without checking
func StraightResponseToStruct(ip_address string) (app.MinerData, error) {
	var result app.MinerData

	response, err := GetAsicInfo(ip_address, "summary")
	if err != nil {
		return result, err
	}

	result, err = ParsingData(response)

	return result, err
}

// Accept string and try parse it to MinerData struct
// using regex.
func ParsingData(data string) (app.MinerData, error) {
	var minerData app.MinerData

	// Searches key and value pairs in data string
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

// func checks is it actually asic by IP
func CheckResponse(response string) error {
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
