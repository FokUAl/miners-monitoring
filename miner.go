package app

type MinerDevice struct {
	Id          int
	MinerType   string
	Shelf       int
	Row         int
	Column      int
	Owner       string
	MinerStatus string
	Coin        string
	IPAddress   string
	MACAddress  string
	Pool        string
}
