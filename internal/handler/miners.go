package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	var devices []app.MinerDevice

	devicesInfo, err := h.services.GetDevicesInfo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s\n", err.Error()))
		return
	}

	deviceResponse := make(chan app.MinerData)
	for _, elem := range devicesInfo {
		var device app.MinerDevice

		device.MinerType = elem.MinerType
		device.IPAddress = elem.IP
		device.Shelf = elem.Shelf
		device.Row = elem.Row
		device.Column = elem.Column
		device.Owner = elem.Owner

		// start goroutune and
		// send result to channel
		go pkg.ResponseToStruct(elem.IP, deviceResponse)

		devices = append(devices, device)
	}

	// reading data from channel
	for i := 0; i < len(devicesInfo); i++ {
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

		isFreeLocation, err := h.services.IsLocationFree(info.Data[i].Shelf, info.Data[i].Row, info.Data[i].Column)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeLocation {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Location isn't free: %d-%d-%d\n",
				info.Data[i].Shelf, info.Data[i].Column, info.Data[i].Row)})
			return
		}

		isFreeIP, err := h.services.IsIPFree(info.Data[i].IP)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeIP {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Device with this IP exists: %s", info.Data[i].IP)})
			return
		}

	}

	for j := 0; j < len(info.Data); j++ {
		var device app.MinerDevice

		device.IPAddress = info.Data[j].IP
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
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s\n", err.Error()))
		return
	}

	deviceResponse := make(chan app.MinerData)
	for _, elem := range devicesInfo {
		var device app.MinerDevice

		device.MinerType = elem.MinerType
		device.IPAddress = elem.IP
		device.Shelf = elem.Shelf
		device.Row = elem.Row
		device.Column = elem.Column
		device.Owner = elem.Owner

		// start goroutune and
		// send result to channel
		go pkg.ResponseToStruct(elem.IP, deviceResponse)

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
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't convert to int: %s\n", err.Error()))
		return
	}
	row, err := strconv.Atoi(query_params["row"][0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't convert to int: %s\n", err.Error()))
		return
	}
	column, err := strconv.Atoi(query_params["column"][0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't convert to int: %s\n", err.Error()))
		return
	}

	miner, err := h.services.GetDeviceByLocation(shelf, column, row)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("get device by location: %s\n", err.Error()))
		return
	}

	channel := make(chan app.MinerData)
	go pkg.ResponseToStruct(miner.IPAddress, channel)
	miner.Characteristics = <-channel

	if miner.Characteristics.MHSav == 0.0 {
		miner.MinerStatus = "offline"
	}

	c.JSON(http.StatusOK, miner)
}

func (h *Handler) UpdateAsicInfo(c *gin.Context) {
	var info app.AddInfo

	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
		return
	}

}
