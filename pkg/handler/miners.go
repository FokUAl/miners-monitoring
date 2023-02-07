package handler

import (
	"fmt"
	"html/template"
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
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	var devices []app.MinerDevice
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
		log.Printf("getHome: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	id := c.MustGet("id").(int)
	user, err := h.services.GetUserByID(id)
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	newInfo := info{
		User:    user,
		Devices: devices,
	}

	err = t.Execute(c.Writer, newInfo)
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getHome: %s", err.Error()))
		return
	}
}

func (h *Handler) getAddMiner(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/add-new.html")
	if err != nil {
		log.Printf("getAddMiner: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getAddMiner: %s", err.Error()))
		return
	}

	notificationText, err := c.Cookie("ErrorContent")
	err = t.Execute(c.Writer, notificationText)

	if err != nil {
		log.Printf("getAddMiner: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getAddMiner: %s", err.Error()))
		return
	}
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
	t, err := template.ParseFiles("./ui/html/grid.html")
	if err != nil {
		log.Printf("minersGrid: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("minersGrid: %s", err.Error()))
		return
	}

	devices, err := h.services.GetAllDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s", err.Error()))
	}

	err = t.Execute(c.Writer, devices)
	if err != nil {
		log.Printf("minersGrid: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("minersGrid: %s", err.Error()))
		return
	}
}

func (h *Handler) getMinerCharacteristics(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/asic.html")
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	shelf, err := strconv.Atoi(c.PostForm("shelf"))
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	row, err := strconv.Atoi(c.PostForm("row"))
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}

	column, err := strconv.Atoi(c.PostForm("column"))
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

	err = t.Execute(c.Writer, device)
	if err != nil {
		log.Printf("getMinerCharacteristics: %s\n", err.Error())
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("getMinerCharacteristics: %s", err.Error()))
		return
	}
}
