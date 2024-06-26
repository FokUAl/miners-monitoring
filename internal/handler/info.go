package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	app "github.com/FokUAl/miners-monitoring"
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

	var devicesAddresses [][]string

	for _, value := range allAddresses {
		response, _ := pkg.GetAsicInfo(value[1], "summary")

		err = pkg.CheckResponse(response)
		if err != nil {
			log.Printf("check response %s: %s", value, err.Error())
			continue
		}
		deviceType := h.services.DefineType(response)
		value = append(value, deviceType)

		isMapped, err := h.services.IsIPMapped(value[1])
		if err != nil {
			log.Printf("check response %s: %s", value, err.Error())
			continue
		}

		if !isMapped {
			devicesAddresses = append(devicesAddresses, value)
		}
	}

	h.services.SaveAvailableAddresses(devicesAddresses)

	c.JSON(http.StatusOK, struct {
		List [][]string
	}{
		List: devicesAddresses,
	})

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

func (h *Handler) SaveMinerData(c *gin.Context, exitChan chan bool) {
	for {
		select {
		case <-exitChan:
			return
		default:
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
				go pkg.ResponseToStruct(elem.IP, device.MinerType, deviceResponse)

				devices = append(devices, device)
			}

			// reading data from channel
			for i := 0; i < len(devicesInfo); i++ {
				responseData := <-deviceResponse
				pkg.UpdateDeviceInfo(&devices, responseData)
			}

			// Saving data to database
			for j := 0; j < len(devices); j++ {
				h.services.SaveMinerData(devices[j].Characteristics, devices[j].IPAddress)
			}

			time.Sleep(59 * time.Second)
		}
	}
}

func (h *Handler) GetKernelLog(c *gin.Context) {
	type Container struct {
		Content string `json:"IP"`
	}
	var tempCont Container
	err := json.NewDecoder(c.Request.Body).Decode(&tempCont)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log := h.services.GetKernelLog(tempCont.Content)
	if log == "" {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, struct {
			KernelLog string
		}{
			KernelLog: log,
		})
	}
}
