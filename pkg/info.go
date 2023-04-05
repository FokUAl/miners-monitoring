package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	app "github.com/FokUAl/miners-monitoring"
)

// Sends request to cgminer interface and returns
// response from it in string type.
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

	var reply bytes.Buffer
	io.Copy(&reply, con)

	if err != nil {
		return "", fmt.Errorf("GetAsicInfo: %s", err.Error())
	}

	return reply.String(), nil
}

// Func for getting data by IP and
// transforming data to MinerData struct
func ResponseToStruct(ip_address string, downstream chan app.MinerData) {
	var result app.MinerData
	result.IP = ip_address
	response, err := GetAsicInfo(ip_address, "stats")
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

	result, err = ParsingDataOld(response)
	if err != nil {
		log.Printf("ResponseToStruct: %s", err.Error())
	}
	result.IP = ip_address
	downstream <- result
}

// Check device characteristics and set miner status
func UpdataDeviceInfo(devices *[]app.MinerDevice, newData app.MinerData) {
	for i := 0; i < len(*devices); i++ {
		if (*devices)[i].IPAddress == newData.IP {
			(*devices)[i].Characteristics = newData
		}
		if (*devices)[i].Characteristics.Temperature == 0.0 {
			(*devices)[i].MinerStatus = "offline"
		} else {
			(*devices)[i].MinerStatus = "online"
		}
	}
}

// Accept string and try parse it to MinerData struct
// using regex. Only for cgminer newer than 4 version.
func ParsingDataNew(data string) (app.MinerData, error) {
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

// Accept string and try parse it to MinerData struct
// using regex. Only for cgminer 1 version.
func ParsingDataOld(data string) (app.MinerData, error) {
	var minerData app.MinerData

	// Searches key and value pairs in data string
	r, err := regexp.Compile(`"[A-Za-z0-9% _]+":"?[0-9A-Za-z:._ -]+"?`)
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

	minerData.PowerMode = "Normal"

	temp_pcb1 := strings.Split(data_map["temp_pcb1"], "-")
	temp_pcb2 := strings.Split(data_map["temp_pcb2"], "-")
	temp_pcb3 := strings.Split(data_map["temp_pcb3"], "-")
	minerData.Temperature, err = Max(temp_pcb1, temp_pcb2, temp_pcb3)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp of boards: %s", err.Error())
	}

	temp_chip1 := strings.Split(data_map["temp_chip1"], "-")
	temp_chip2 := strings.Split(data_map["temp_chip2"], "-")
	temp_chip3 := strings.Split(data_map["temp_chip3"], "-")
	chip_temp_stats, err := ChipTempStats(temp_chip1, temp_chip2, temp_chip3)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp of chips: %s", err.Error())
	}

	minerData.ChipTempMax = chip_temp_stats[0]
	minerData.ChipTempMin = chip_temp_stats[1]
	minerData.ChipTempAvg = chip_temp_stats[2]

	minerData.MHSav, err = strconv.ParseFloat(data_map["GHS av"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse GHS av: %s", err.Error())
	}

	minerData.FanSpeedIn, err = strconv.ParseInt(data_map["fan1"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed In: %s", err.Error())
	}

	minerData.FanSpeedOut, err = strconv.ParseInt(data_map["fan3"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed Out: %s", err.Error())
	}

	minerData.Elapsed, err = strconv.ParseInt(data_map["Elapsed"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Elapsed: %s", err.Error())
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

// finds the maximum number in arrays.
func Max(arr1 []string, arr2 []string, arr3 []string) (float64, error) {
	max := 0.0
	arr1 = append(arr1, arr2...)
	arr1 = append(arr1, arr3...)

	nums, err := ParseFloatArr(arr1)
	if err != nil {
		return 0.0, fmt.Errorf("Max: %s", err.Error())
	}

	for _, elem := range nums {
		if elem > max {
			max = elem
		}
	}

	return max, nil
}

// finds max temp, min temp, average temp for chips.
func ChipTempStats(arr1 []string, arr2 []string, arr3 []string) ([]float64, error) {
	// combining arrays
	arr1 = append(arr1, arr2...)
	arr1 = append(arr1, arr3...)

	nums, err := ParseFloatArr(arr1)
	if err != nil {
		return nil, fmt.Errorf("ChipTempStats: %s", err.Error())
	}

	max := 0.0
	min := 100.0
	sum := 0.0
	for _, elem := range nums {
		if elem < min {
			min = elem
		}
		if elem > max {
			max = elem
		}
		sum += elem
	}

	result := []float64{max, min, sum / 12}

	return result, nil
}

// parse array of string to array of float
func ParseFloatArr(arr []string) ([]float64, error) {
	var result []float64
	for i := 0; i < len(arr); i++ {
		num, err := strconv.ParseFloat(arr[i], 64)
		if err != nil {
			return nil, err
		}

		result = append(result, num)
	}

	return result, nil
}
