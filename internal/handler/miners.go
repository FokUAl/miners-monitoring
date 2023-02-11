package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	app "github.com/FokUAl/miners-monitoring"
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

func (h *Handler) addMiner(c *gin.Context) {
	isIP := c.PostForm("connection") == "ip"

	connections := c.PostFormArray("ip/mac")
	shelfData := c.PostFormArray("shelf")
	rowData := c.PostFormArray("row")
	columnData := c.PostFormArray("column")

	// info about location
	locInfo := [][]string{shelfData, rowData, columnData}

	err := h.services.Miner.AddDevices(c.PostForm("model"), isIP, connections, locInfo)

	if err != nil {
		c.SetCookie("ErrorContent", err.Error(), 10, "/add", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/add")
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func (h *Handler) minersGrid(c *gin.Context) {

	devices, err := h.services.GetAllDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s", err.Error()))
	}

	newInfo := info{
		Devices: devices,
	}

	c.JSON(http.StatusOK, newInfo)
}

func (h *Handler) getMinerCharacteristics(c *gin.Context) {
	values := c.Request.URL.Query()

	shelf, err := strconv.Atoi(values["shelf"][0])
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	row, err := strconv.Atoi(values["row"][0])
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	column, err := strconv.Atoi(values["column"][0])
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	device, err := h.services.GetDeviceByLocation(shelf, column, row)
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, device)
}
