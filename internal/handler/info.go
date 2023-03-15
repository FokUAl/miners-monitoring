package handler

import (
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

	var devicesAdresses map[string]string = make(map[string]string)

	for key, value := range allAddresses {
		response, _ := pkg.GetAsicInfo(value, "summary")

		err = pkg.CheckResponse(response)
		if err != nil {
			log.Printf("check response %s: %s", value, err.Error())
			continue
		}
		devicesAdresses[key] = value
	}

	type IPDevices struct {
		List map[string]string
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

// func SaveMinerData(h *Handler) {
// 	for {
// 		time.Sleep(10 * time.Second)
// 		h.services.
// 	}
// }
