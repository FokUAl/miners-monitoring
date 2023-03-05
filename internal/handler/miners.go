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
	Notification  string
	User          app.User
	FormedDevices [][]app.MinerDevice
	Devices       []app.MinerDevice
}

func (h *Handler) getHome(c *gin.Context) {
	var devices []app.MinerDevice
	var err error
	filter_category := c.PostForm("category")

	switch filter_category {
	case "Miner Type":
		miner_type := c.PostForm("target")
		devices, err = h.services.GetDevicesByType(miner_type)
	case "Status":
		status := c.PostForm("target")
		devices, err = h.services.GetDevicesByStatus(status)
	case "Coin":
		coin := c.PostForm("target")
		devices, err = h.services.GetDevicesByCoin(coin)
	default:
		devices, err = h.services.GetAllDevices()
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	id := c.MustGet(userCtx).(int)
	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	reformedDevices, err := h.services.Transform(devices)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	newInfo := info{
		User:          user,
		FormedDevices: reformedDevices,
	}

	c.JSON(http.StatusOK, newInfo)
}

func (h *Handler) addMiner(c *gin.Context) {
	type MappingInfo struct {
		Data []app.AddInfo
	}

	var info MappingInfo

	err := json.NewDecoder(c.Request.Body).Decode(&info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
		return
	}

	for i := 0; i < len(info.Data); i++ {
		var device app.MinerDevice

		device.IPAddress = info.Data[i].IP
		device.Shelf = info.Data[i].Shelf
		device.Row = info.Data[i].Row
		device.Column = info.Data[i].Column
		device.Owner = info.Data[i].Owner

		isFreeLocation, err := h.services.IsLocationFree(device.Shelf, device.Row, device.Column)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeLocation {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Location isn't free: %d-%d-%d\n", device.Shelf, device.Column, device.Row)})
			return
		}

		isFreeIP, err := h.services.IsIPFree(device.IPAddress)
		if err != nil && err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
		if !isFreeIP {
			c.JSON(http.StatusBadRequest, Notification{Message: fmt.Sprintf("Device with this IP exists: %s", device.IPAddress)})
			return
		}

		err = h.services.AddNew(device)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}
	}

	err = h.services.MappDevices(info.Data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
		return
	}

}

// Heat map that explain location and temperature of device
func (h *Handler) minersGrid(c *gin.Context) {

	devices, err := h.services.GetAllDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s", err.Error()))
		return
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

	info, err := pkg.GetAsicInfo(miner.IPAddress, "summary")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with sending command: %s\n", err.Error()))
		return
	}

	err = h.services.CheckResponse(info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't check response: %s\n", err.Error()))
		return
	}

	strct, err := h.services.Info.ParsingData(info)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't parse data: %s\n", err.Error()))
		return
	}

	c.JSON(http.StatusOK, strct)
}
