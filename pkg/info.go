package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	app "github.com/FokUAl/miners-monitoring"
)

// Uses python script to send request to cgminer interface
// and return response from it as string.
func GetAsicInfo(ip string, command string) (string, error) {
	d := net.Dialer{Timeout: time.Second}
	con, err := d.Dial("tcp", ip+":4028")
	if err != nil {
		return "", fmt.Errorf("GetAsicInfo: %s", err.Error())
	}

	defer con.Close()

	payload := make(map[string]string)
	payload["command"] = command
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("GetAsicInfo: %s", err.Error())
	}

	_, err = con.Write(jsonPayload)
	if err != nil {
		return "", fmt.Errorf("GetAsicInfo: %s", err.Error())
	}

	reply := make([]byte, 2048)
	_, err = con.Read(reply)

	if err != nil {
		return "", fmt.Errorf("GetAsicInfo: %s", err.Error())
	}

	return string(reply), nil
}

// Func for getting data by IP and
// transforming data to MinerData struct
func ResponseToStruct(ip_address string, downstream chan app.MinerData) {
	var result app.MinerData
	result.IP = ip_address
	response, err := GetAsicInfo(ip_address, "summary")
	if err != nil {
		log.Printf("ResponseToStruct: %s", err.Error())
		downstream <- result
		return
	}

	err = CheckResponse(response)
	if err != nil {
		log.Printf("ResponseToStruct: %s", err.Error())
		downstream <- result
		return
	}

	result, err = ParsingData(response)
	if err != nil {
		log.Printf("ResponseToStruct: %s", err.Error())
	}
	result.IP = ip_address
	downstream <- result
}

func UpdataDeviceInfo(devices *[]app.MinerDevice, newData app.MinerData) {
	for i := 0; i < len(*devices); i++ {
		if (*devices)[i].IPAddress == newData.IP {
			(*devices)[i].Characteristics = newData
		}
		if (*devices)[i].Characteristics.MHSav == 0.0 {
			(*devices)[i].MinerStatus = "offline"
		} else {
			(*devices)[i].MinerStatus = "online"
		}
	}
}

// Accept string and try parse it to MinerData struct
// using regex.
func ParsingData(data string) (app.MinerData, error) {
	var minerData app.MinerData

	// Searches key and value pairs in data string
	r, err := regexp.Compile(`"[A-Za-z0-9% ]+":("?[0-9A-Za-z:._ -]+"?)`)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't compile regexp: %s", err.Error())
	}

	arr := r.FindAllString(data, -1)
	data_map := make(map[string]string)
	for _, val := range arr {
		keyvalue := strings.Split(val, "\":")
		key := strings.Trim(keyvalue[0], "\"")
		value := strings.Trim(keyvalue[1], "\"")
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
	if len(response) < 30 || response[13:30] == "'STATUS': 'error'" {
		return fmt.Errorf("CheckResponse: error response: [%s]", response)
	}

	return nil
}
