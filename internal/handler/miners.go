package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
	Shelf  int    `json:"shelf"`
	Column int    `json:"column"`
	Row    int    `json:"row"`
	Owner  int    `json:"owner"`
}

type FrontResponse struct {
	Objects []AddInfo
}

func (h *Handler) addMiner(c *gin.Context) {
	var frontData FrontResponse
	dec := json.NewDecoder(c.Request.Body)

	// t, err := dec.Token()
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError,
	// 		fmt.Sprintf("addMiner: %s", err.Error()))
	// 	return
	// }
	// log.Printf("%T: %v\n", t, t)

	count := 0
	for dec.More() {
		count += 1
		var info AddInfo
		err := dec.Decode(&info)
		log.Printf("%d : %v\n", count, info)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError,
				fmt.Sprintf("addMiner: %s", err.Error()))
			return
		}

		frontData.Objects = append(frontData.Objects, info)
	}

	// t, err = dec.Token()
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError,
	// 		fmt.Sprintf("addMiner: %s", err.Error()))
	// 	return
	// }
	// log.Printf("%T: %v\n", t, t)

	//log.Println(frontData)
	//err := h.services.Miner.AddDevices(c.PostForm("model"), isIP, connections, locInfo)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, nil)
	// } else {
	// 	c.JSON(http.StatusOK, nil)
	// }
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
