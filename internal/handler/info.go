package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FokUAl/miners-monitoring/pkg"
	"github.com/gin-gonic/gin"
)

// Searches for the addresses of all devices in the local network
// and determines the ASIC among them.
// Result is sent to front.
func (h *Handler) FindDeviceIP(c *gin.Context) {
	allAddresses, err := h.services.PingDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError,
			fmt.Sprintf("error with ping devices: %s\n", err.Error()))
		return
	}

	var devicesAdresses [][]string

	for _, value := range allAddresses {
		response, _ := pkg.GetAsicInfo(value[1], "summary")

		err = pkg.CheckResponse(response)
		if err != nil {
			log.Printf("check response %s: %s", value, err.Error())
			continue
		}
		devicesAdresses = append(devicesAdresses, value)
	}

	type IPDevices struct {
		List [][]string
	}

	var IP IPDevices
	IP.List = devicesAdresses

	c.JSON(http.StatusOK, IP)

}

func (h *Handler) GetUserInfo(c *gin.Context) {
	id := c.MustGet(userCtx).(int)
	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) CommentDevice(c *gin.Context) {
	type comment struct {
		Content string `json:"comment"`
		Address string `json:"IP"`
	}

	var input comment
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("CommentDevice: %s", err.Error()))
		return
	}

	id := c.MustGet(userCtx).(int)
	user, err := h.services.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

	err = h.services.Info.Comment(input.Address, user.Username, input.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("getHome: %s", err.Error()))
		return
	}

}
