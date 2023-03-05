package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FokUAl/miners-monitoring/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FindDeviceIP(c *gin.Context) {
	ipArr, err := h.services.PingDevices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with ping devices: %s\n", err.Error()))
		return
	}

	var devicesIP []string
	for i := 0; i < len(ipArr); i++ {
		response, err := pkg.GetAsicInfo(ipArr[i], "summary")
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with get info from devices: %s\n", err.Error()))
			return
		}

		err = h.services.CheckResponse(response)
		if err != nil {
			log.Printf("check response %s: %s", ipArr[i], err.Error())
			continue
		}
		devicesIP = append(devicesIP, ipArr[i])
	}

	type IPDevices struct {
		List []string
	}

	var IP IPDevices
	IP.List = devicesIP

	c.JSON(http.StatusOK, IP)

}

// func SaveMinerData(h *Handler) {
// 	for {
// 		time.Sleep(10 * time.Second)
// 		h.services.
// 	}
// }
