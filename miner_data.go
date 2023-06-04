package app

import "time"

type MinerData struct {
	IP          string
	Elapsed     int64
	THSav       float64
	Temperature float64
	FanSpeed1   int64
	FanSpeed2   int64
	FanSpeed3   int64
	FanSpeed4   int64
	PowerMode   string
	MAC         string
	ChipTempMin float64
	ChipTempMax float64
	ChipTempAvg float64
	Share 		int64
	Time        time.Time
}

type AddInfo struct {
	IP        string `json:"IP"`
	MAC       string `json:"MAC"`
	MinerType string `json:"minerType"`
	Shelf     int    `json:"shelf,string"`
	Column    int    `json:"column,string"`
	Row       int    `json:"row,string"`
	Owner     string `json:"owner"`
}

type Comment struct {
	CreationDate time.Time
	Username     string
	Content      string
}
