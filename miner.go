package app

type MinerDevice struct {
	Id              int
	MinerType       string
	Shelf           int
	Row             int
	Column          int
	Owner           string
	MinerStatus     string
	IPAddress       string
	MACAddress      string
	Characteristics MinerData
}
