package app

type MinerData struct {
	IP          string
	Elapsed     int64
	MHSav       float64
	Temperature float64
	FanSpeedIn  int64
	FanSpeedOut int64
	PowerMode   string
	MAC         string
	ChipTempMin float64
	ChipTempMax float64
	ChipTempAvg float64
}

type AddInfo struct {
	IP        string `json:"IP"`
	MinerType string //`json:"miner-type"`
	Shelf     int    `json:"shelf,string"`
	Column    int    `json:"column,string"`
	Row       int    `json:"row,string"`
	Owner     string `json:"owner"`
}
