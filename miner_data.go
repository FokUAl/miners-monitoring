package app

type MinerData struct {
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
