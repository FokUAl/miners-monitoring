package service

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	app "github.com/FokUAl/miners-monitoring"
	"github.com/FokUAl/miners-monitoring/internal/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{
		repo: repo,
	}
}

// Ping all IP adresses in specific range
// and determine alive devices from them.
// Return result as array of strings
func (s *InfoService) PingDevices() ([][]string, error) {
	var result [][]string

	// Pinging in specific range for update ARP cache
	cmd := exec.Command("fping", "-g", "-a", "-i 1", "-r 0", "192.168.0.0/24")
	cmd.Run()
	// Determines alive devices
	cmd = exec.Command("bash", "-c", "arp -an | grep 'ether'")

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("pingDevices: %s", err.Error())
	}

	r, err := regexp.Compile(`\(.+\)`)
	if err != nil {
		return nil, fmt.Errorf("pingDevices: can't compile regexp: %s", err.Error())
	}

	allIP := r.FindAllString(string(out), -1)
	for ind := 0; ind < len(allIP); ind++ {
		allIP[ind] = strings.Trim(allIP[ind], "()")
	}

	r, err = regexp.Compile(`([a-f0-9]{2}[:]){5}([a-f0-9]{2})`)
	if err != nil {
		return nil, fmt.Errorf("pingDevices: can't compile regexp: %s", err.Error())
	}

	allMAC := r.FindAllString(string(out), -1)
	for ind := 0; ind < len(allMAC); ind++ {
		var arr []string = []string{allMAC[ind], allIP[ind]}
		result = append(result, arr)
	}

	return result, nil
}

func (s *InfoService) SaveMinerData(data app.MinerData, ip_address string) error {
	return s.repo.SaveMinerData(data, ip_address)
}

func (s *InfoService) Transform(devices []app.MinerDevice) (map[string][]app.MinerData, error) {
	var result map[string][]app.MinerData = make(map[string][]app.MinerData)
	if len(devices) == 0 {
		return nil, fmt.Errorf("transform: argument array is empty")
	}

	for _, dev := range devices {
		_, ok := result[dev.Owner]
		if ok {
			result[dev.Owner] = append(result[dev.Owner], dev.Characteristics)
		} else {
			result[dev.Owner] = []app.MinerData{dev.Characteristics}
		}

	}

	return result, nil
}

func (s *InfoService) GetCharacteristicsHistory(device_ip string) ([]app.MinerData, error) {
	return s.repo.GetCharacteristicsHistory(device_ip)
}

func (s *InfoService) DetermineIP(mac_address string) string {
	return s.repo.DetermineIP(mac_address)
}

func (s *InfoService) DetermineMAC(ip_address string) string {
	return s.repo.DetermineMAC(ip_address)
}

func (s *InfoService) SaveAvailableAddresses(list [][]string) error {
	return s.repo.SaveAvailableAddresses(list)
}

func (s *InfoService) IsIPMapped(ip string) (bool, error) {
	return s.repo.IsIPMapped(ip)
}
