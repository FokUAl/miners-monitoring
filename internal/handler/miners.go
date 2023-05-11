package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/pkg"
	"github.com/gin-gonic/gin"
)

type info struct {
	Notification string
	FormedData   map[string][]app.MinerData
	Devices      []app.MinerDevice
}

type MappingInfo struct {
	Data []app.AddInfo
}

func (h *Handler) getHome(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.Trim(token, "\"")
	id, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var devices []app.MinerDevice

	devicesInfo, err := h.services.GetDevicesInfo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s\n", err.Error()))
		return
	}

	deviceResponse := make(chan app.MinerData)
	for _, elem := range devicesInfo {
		var device app.MinerDevice

		device.IPAddress = elem.IP

		if user.Role == "User" &&
			elem.Owner != user.Username {
			continue
		}

		device.MinerType = elem.MinerType
		device.Shelf = elem.Shelf
		device.Row = elem.Row
		device.Column = elem.Column
		device.Owner = elem.Owner

		// start goroutune and
		// send result to channel
		go pkg.ResponseToStruct(device.IPAddress, deviceResponse)

		devices = append(devices, device)
	}

	// reading data from channel
	for i := 0; i < len(devices); i++ {
		responseData := <-deviceResponse
		pkg.UpdataDeviceInfo(&devices, responseData)
	}

	var formedDeviceData map[string][]app.MinerData = make(map[string][]app.MinerData)
	if len(devices) != 0 {
		formedDeviceData, err = h.services.Transform(devices)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
			return
		}
	}

	newInfo := info{
		Devices:    devices,
		FormedData: formedDeviceData,
	}

	c.JSON(http.StatusOK, newInfo)
}

func (h *Handler) addMiner(c *gin.Context) {
	var info MappingInfo

	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
		return
	}

	for i := 0; i < len(info.Data); i++ {
		info.Data[i].MAC = strings.ToLower(info.Data[i].MAC)
		isFreeLocation, err := h.services.IsLocationFree(info.Data[i].Shelf, info.Data[i].Row, info.Data[i].Column)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeLocation {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Location isn't free: %d-%d-%d\n",
				info.Data[i].Shelf, info.Data[i].Column, info.Data[i].Row)})
			log.Printf("addMiner: Location isn't free: %d-%d-%d\n", info.Data[i].Shelf, info.Data[i].Column, info.Data[i].Row)
			return
		}

		isFreeAddress, err := h.services.IsAddressFree(info.Data[i].IP, info.Data[i].MAC)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeAddress {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Device with this address exists: %s", info.Data[i].IP)})
			return
		}

	}

	for j := 0; j < len(info.Data); j++ {
		var device app.MinerDevice

		if info.Data[j].MAC == "" {
			device.IPAddress = info.Data[j].IP
			device.MACAddress = h.services.DetermineMAC(device.IPAddress)
		} else {
			device.MACAddress = info.Data[j].MAC
			device.IPAddress = h.services.DetermineIP(device.MACAddress)
		}

		device.Shelf = info.Data[j].Shelf
		device.Row = info.Data[j].Row
		device.Column = info.Data[j].Column
		device.Owner = info.Data[j].Owner
		device.MinerType = info.Data[j].MinerType

		err = h.services.AddNew(device)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
	}
}

// Heat map that explain location and temperature of device
func (h *Handler) minersGrid(c *gin.Context) {

	var devices []app.MinerDevice

	devicesInfo, err := h.services.GetDevicesInfo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s\n", err.Error()))
		return
	}

	deviceResponse := make(chan app.MinerData)
	for _, elem := range devicesInfo {
		var device app.MinerDevice

		address := elem.IP

		// isIP, err := pkg.IsIP(address)
		// if err != nil {
		// 	newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s\n", err.Error()))
		// 	return
		// }

		device.MinerType = elem.MinerType
		device.Shelf = elem.Shelf
		device.Row = elem.Row
		device.Column = elem.Column
		device.Owner = elem.Owner

		// start goroutune and
		// send result to channel
		go pkg.ResponseToStruct(address, deviceResponse)

		devices = append(devices, device)
	}

	// reading data from channel
	for i := 0; i < len(devicesInfo); i++ {
		responseData := <-deviceResponse
		pkg.UpdataDeviceInfo(&devices, responseData)
	}

	newInfo := info{
		Devices: devices,
	}

	c.JSON(http.StatusOK, newInfo)
}

// Find device IP by location and get characteristics like temperature,
// hashrate and other.
func (h *Handler) getMinerCharacteristics(c *gin.Context) {
	query_params := c.Request.URL.Query()
	shelf, err := strconv.Atoi(query_params["shelf"][0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinerCharacteristics: can't convert to int: %s\n", err.Error()))
		return
	}
	row, err := strconv.Atoi(query_params["row"][0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinerCharacteristics: can't convert to int: %s\n", err.Error()))
		return
	}
	column, err := strconv.Atoi(query_params["column"][0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinerCharacteristics: can't convert to int: %s\n", err.Error()))
		return
	}

	miner, err := h.services.GetDeviceByLocation(shelf, column, row)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinerCharacteristics: %s\n", err.Error()))
		return
	}

	channel := make(chan app.MinerData)
	go pkg.ResponseToStruct(miner.IPAddress, channel)
	miner.Characteristics = <-channel

	if miner.Characteristics.THSav == 0.0 {
		miner.MinerStatus = "offline"
	}

	comments, err := h.services.GetCommentsHistory(miner.IPAddress)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinersCharacteristics: %s\n", err.Error()))
		return
	}

	characteristicsHistory, err := h.services.GetCharacteristicsHistory(miner.IPAddress)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getMinersCharacteristics: %s\n", err.Error()))
		return
	}

	c.JSON(http.StatusOK, struct {
		Miner                  app.MinerDevice
		Comments               []app.Comment
		CharacteristicsHistory []app.MinerData
	}{
		Miner:                  miner,
		Comments:               comments,
		CharacteristicsHistory: characteristicsHistory,
	})
}

// Uses when we change asic info(miner type, owner, location)
func (h *Handler) UpdateAsicInfo(c *gin.Context) {
	var info app.AddInfo

	// get new info from front
	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("UpdateAsicInfo: %s", err.Error()))
		return
	}

	// getting current device info
	device, err := h.services.GetDeviceByLocation(info.Shelf, info.Column, info.Row)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("UpdateAsicInfo: %s", err.Error()))
		return
	}

	var infoHolder []app.AddInfo
	infoHolder = append(infoHolder, info)

	var address string
	if info.MAC == "" {
		address = device.IPAddress
	} else {
		address = device.MACAddress
	}

	// checking if new location is free
	isLocFree, err := h.services.IsLocationFree(info.Shelf, info.Row, info.Column)
	if err != nil && err != sql.ErrNoRows {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("UpdateAsicInfo: %s", err.Error()))
		return
		// check if we write same address
	} else if !isLocFree && info.IP != address {
		c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Location isn't free: %d-%d-%d\n",
			info.Shelf, info.Column, info.Row)})
		return
	}

	// update info in database
	err = h.services.MappDevices(infoHolder)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("UpdateAsicInfo: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteDevice(c *gin.Context) {
	type inputJson struct {
		IP string `json:"IP"`
	}

	var inputInfo inputJson

	err := json.NewDecoder(c.Request.Body).Decode(&inputInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteDevice: %s", err.Error()))
		return
	}

	err = h.services.DeleteDevice(inputInfo.IP)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("DeleteDevice: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}
