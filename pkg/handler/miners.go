package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

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

	devices, err := h.services.GetAllDevices()
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

	err = t.Execute(c.Writer, nil)
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
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("addMiner: %s", err.Error()))
	}

	c.Redirect(http.StatusSeeOther, "/")
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
