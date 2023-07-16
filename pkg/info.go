package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	app "github.com/FokUAl/miners-monitoring"
)

type ShareStruct struct {
	IP    string
	Value int64
}

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
func ResponseToStruct(ip_address string, minerType string, downstream chan app.MinerData) {
	var result app.MinerData

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

	minerType = strings.Split(minerType, " ")[0]
	if minerType == "" {
		log.Println("ResponseToStruct: can't determine type of miner")
		downstream <- result
		return
	}

	//need to check
	switch minerType {
	case "Whatsminer":
		//right
		result, err = ParsingDataNew(response)
		if err != nil {
			log.Printf("ResponseToStruct: %s", err.Error())
		}
	case "Antminer":
		response, err := GetAsicInfo(ip_address, "stats")
		if err != nil {
			log.Printf("ResponseToStruct: %s", err.Error())
		}

		result, err = ParsingDataOld(response)
		if err != nil {
			log.Printf("ResponseToStruct: %s", err.Error())
		}

	case "Avalon":
		response, err := GetAsicInfo(ip_address, "stats")
		if err != nil {
			log.Printf("ResponseToStruct: %s", err.Error())
		}

		result, err = ParsingDataMiddle(response)
		if err != nil {
			log.Printf("ResponseToStruct: %s", err.Error())
		}
	}

	result.IP = ip_address
	downstream <- result
}

func GetShare(IP string, share chan ShareStruct) {
	response, err := GetAsicInfo(IP, "summary")
	if err != nil {
		share <- ShareStruct{IP: err.Error(), Value: 0}
		return
	}
	shareInt, err := ShareParsing(response)
	if err != nil {
		share <- ShareStruct{IP: err.Error(), Value: 0}
		return
	}
	share <- ShareStruct{IP, shareInt}
}

func UpdateDeviceShare(devices *[]app.MinerDevice, share ShareStruct) {
	for i := 0; i < len(*devices); i++ {
		if (*devices)[i].IPAddress == share.IP {
			(*devices)[i].Characteristics.Share = share.Value
		}
	}
}

func ShareParsing(data string) (int64, error) {
	// Searches key and value pairs in data string
	r, err := regexp.Compile(`"[A-Za-z0-9% ]+":("?[0-9A-Za-z:._ -]+"?)`)
	if err != nil {
		return 0, fmt.Errorf("can't compile regexp: %s", err.Error())
	}

	arr := r.FindAllString(data, -1)
	data_map := make(map[string]string)
	for _, val := range arr {
		keyvalue := strings.Split(val, "\":")
		key := strings.Trim(keyvalue[0], "\"")
		value := strings.Trim(keyvalue[1], "\"")
		data_map[key] = value
	}

	share, err := strconv.ParseInt(data_map["Accepted"], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("can't parse Accepted: %s", err.Error())
	}
	return share, nil
}

// Check device characteristics and set miner status
func UpdateDeviceInfo(devices *[]app.MinerDevice, newData app.MinerData) {
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

	var ok bool
	minerData.MAC, ok = data_map["MAC"]
	if !ok {
		return app.MinerData{}, fmt.Errorf("can't parse temperature MAC")
	}

	minerData.PowerMode = data_map["Power Mode"]
	minerData.Temperature, err = strconv.ParseFloat(data_map["Temperature"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temperature: %s", err.Error())
	}

	minerData.THSav, err = strconv.ParseFloat(data_map["MHS av"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse MHS av: %s", err.Error())
	}

	minerData.THSav = math.Round(minerData.THSav / 10000)

	minerData.FanSpeed1, err = strconv.ParseInt(data_map["Fan Speed In"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse MHS av: %s", err.Error())
	}

	minerData.FanSpeed2, err = strconv.ParseInt(data_map["Fan Speed Out"], 10, 64)
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

func ParsingDataMiddle(data string) (app.MinerData, error) {
	var minerData app.MinerData

	r, err := regexp.Compile(`([A-Za-z0-9]+\[[0-9. ]+\])`)
	if err != nil {
		return minerData, fmt.Errorf("can't compile regexp: %s", err.Error())
	}

	arr := r.FindAllString(data, -1)
	data_map := make(map[string]string)

	for _, val := range arr {
		keyvalue := strings.Split(val, "[")
		key := keyvalue[0]
		value := strings.Trim(keyvalue[1], "]")
		data_map[key] = value
	}

	workmode, ok := data_map["WORKMODE"]
	if !ok {
		return app.MinerData{}, fmt.Errorf("can't parse workmode")
	}

	if workmode == "0" {
		minerData.PowerMode = "Normal"
	} else {
		minerData.PowerMode = "High"
	}

	minerData.Elapsed, err = strconv.ParseInt(data_map["Elapsed"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse elapsed time: %s", err.Error())
	}

	minerData.Temperature, err = strconv.ParseFloat(data_map["TAvg"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temperature: %s", err.Error())
	}

	chipTempsStr := strings.Split(data_map["MTavg"], " ")
	temp_chip1, err := strconv.ParseFloat(chipTempsStr[0], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp_chip_1: %s", err.Error())
	}

	temp_chip2, err := strconv.ParseFloat(chipTempsStr[1], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp_chip_2: %s", err.Error())
	}

	temp_chip3, err := strconv.ParseFloat(chipTempsStr[2], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp_chip_2: %s", err.Error())
	}

	minerData.ChipTempMax = Max3(temp_chip1, temp_chip2, temp_chip3)
	minerData.ChipTempMin = Min3(temp_chip1, temp_chip2, temp_chip3)
	minerData.ChipTempAvg = Avg3(temp_chip1, temp_chip2, temp_chip3)

	minerData.FanSpeed1, err = strconv.ParseInt(data_map["Fan1"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse FanSpeed1: %s", err.Error())
	}

	minerData.FanSpeed2, err = strconv.ParseInt(data_map["Fan2"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse FanSpeed2: %s", err.Error())
	}

	minerData.FanSpeed3, err = strconv.ParseInt(data_map["Fan3"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse FanSpeed3: %s", err.Error())
	}

	minerData.FanSpeed4, err = strconv.ParseInt(data_map["Fan4"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse FanSpeed4: %s", err.Error())
	}

	GHSavg, err := strconv.ParseFloat(data_map["GHSavg"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse GHSavg: %s", err.Error())
	}

	minerData.THSav = math.Round(GHSavg / 1000)

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

	_, ok := data_map["temp_pcb1"]
	if !ok {
		return app.MinerData{}, fmt.Errorf("can't parse temp_pcb1")
	}

	temp_pcb := strings.Split(data_map["temp_pcb1"], "-")
	temp_pcb2 := strings.Split(data_map["temp_pcb2"], "-")
	temp_pcb3 := strings.Split(data_map["temp_pcb3"], "-")

	temp_pcb = append(temp_pcb, temp_pcb2...)
	temp_pcb = append(temp_pcb, temp_pcb3...)

	temp_pcbNum, err := ParseFloatArr(temp_pcb)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse temp of chips: %s", err.Error())
	}

	minerData.Temperature = Max(temp_pcbNum)

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

	GHSav, err := strconv.ParseFloat(data_map["GHS av"], 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse GHS av: %s", err.Error())
	}
	minerData.THSav = math.Round(GHSav / 1000)

	minerData.FanSpeed1, err = strconv.ParseInt(data_map["fan1"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed 1: %s", err.Error())
	}

	minerData.FanSpeed2, err = strconv.ParseInt(data_map["fan2"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed 2: %s", err.Error())
	}
	minerData.FanSpeed3, err = strconv.ParseInt(data_map["fan3"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed 3: %s", err.Error())
	}

	minerData.FanSpeed4, err = strconv.ParseInt(data_map["fan4"], 10, 64)
	if err != nil {
		return app.MinerData{}, fmt.Errorf("can't parse Fan Speed 4: %s", err.Error())
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

// finds the maximum number in array.
func Max(arr []float64) float64 {
	max := 0.0
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}

	return max
}

// finds the minimum number in arrays.
func Min(arr []float64) float64 {
	min := 10000.0
	for _, elem := range arr {
		if elem < min {
			min = elem
		}
	}

	return min
}

func Max3(num1 float64, num2 float64, num3 float64) float64 {
	if num1 >= num2 && num1 >= num3 {
		return num1
	} else if num2 >= num1 && num2 >= num3 {
		return num2
	} else {
		return num3
	}
}

func Avg3(num1 float64, num2 float64, num3 float64) float64 {
	if num1 <= num2 && num1 >= num3 ||
		num1 >= num2 && num1 <= num3 {
		return num1
	} else if num2 <= num1 && num2 >= num3 ||
		num2 >= num1 && num2 <= num3 {
		return num2
	} else {
		return num3
	}
}

func Min3(num1 float64, num2 float64, num3 float64) float64 {
	if num1 <= num2 && num1 <= num3 {
		return num1
	} else if num2 <= num1 && num2 <= num3 {
		return num2
	} else {
		return num3
	}
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

	max := Max(nums)
	min := Min(nums)
	sum := 0.0
	for _, elem := range nums {
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

func IsIP(address string) (bool, error) {
	if address == " " {
		return false, nil
	}
	reg, err := regexp.Compile(`^([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})`)
	if err != nil {
		return false, err
	}

	res := reg.MatchString(address)

	return res, nil
}
