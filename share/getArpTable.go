package share

import (
	"bufio"
	"os"
	"strings"
)

const (
	f_IPAddr uint8 = iota
	f_HWType
	f_Flags
	f_HWAddr
	f_Mask
	f_Device
)

type ArpTableItem struct {
	IPAddr string
	HWType string
	Flags  string
	HWAddr string
	Mask   string
	Device string
}

type ArpTable []ArpTableItem

// GetArpTable
// Read arp table from Linux arp table
func GetArpTable() (ArpTable, error) {
	line := []ArpTableItem{}

	f, err := os.Open("/proc/net/arp")

	if err != nil {
		return line, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()

	for s.Scan() {
		fields := strings.Fields(s.Text())
		line = append(line, ArpTableItem{
			IPAddr: fields[f_IPAddr],
			HWType: fields[f_HWType],
			Flags:  fields[f_Flags],
			HWAddr: fields[f_HWAddr],
			Mask:   fields[f_Mask],
			Device: fields[f_Device],
		})
	}

	return line, nil
}
