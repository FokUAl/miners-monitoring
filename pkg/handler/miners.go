package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getHome(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
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
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
	}

	err = t.Execute(c.Writer, devices)
	if err != nil {
		log.Printf("getHome: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (h *Handler) getAddMiner(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/add-new.html")
	if err != nil {
		log.Printf("getAddMiner: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	notificationText, err := c.Cookie("ErrorContent")
	err = t.Execute(c.Writer, notificationText)

	if err != nil {
		log.Printf("getAddMiner: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
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
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
	// }
	if err != nil {
		c.SetCookie("ErrorContent", err.Error(), 10, "/add", "localhost", false, true)
	}

	c.Redirect(http.StatusSeeOther, "/add")
}

func (h *Handler) minersGrid(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/grid.html")
	if err != nil {
		log.Printf("minersGrid: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	devices, err := h.services.GetAllDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("minersGrid: %s", err.Error()))
	}

	err = t.Execute(c.Writer, devices)
	if err != nil {
		log.Printf("minersGrid: %s\n", err.Error())
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
