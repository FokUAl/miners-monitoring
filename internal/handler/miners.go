package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/pkg"
	"github.com/gin-gonic/gin"
)

type info struct {
	Notification string
	User         app.User
	Devices      []app.MinerDevice
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

	newInfo := info{
		User:    user,
		Devices: devices,
	}

	c.JSON(http.StatusOK, newInfo)
}

func (h *Handler) getAddMiner(c *gin.Context) {
	notificationText, err := c.Cookie("ErrorContent")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getAddMiner: %s", err.Error()))
		return
	}

	newInfo := info{
		Notification: notificationText,
	}
	c.JSON(http.StatusOK, newInfo)
}

type AddInfo struct {
	IP     string `json:"IP"`
	Shelf  int    `json:"shelf,string"`
	Column int    `json:"column,string"`
	Row    int    `json:"row,string"`
	Owner  string `json:"owner"`
}

func (h *Handler) addMiner(c *gin.Context) {
	var mappingInfo []AddInfo

	err := json.NewDecoder(c.Request.Body).Decode(&mappingInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
		return
	}

}

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

func (h *Handler) getMinerCharacteristics(c *gin.Context) {
	info, err := pkg.GetAsicInfo("192.168.0.1", "summary")
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
