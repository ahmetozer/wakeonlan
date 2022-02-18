package share

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	fieldIPAddr uint8 = iota
	fieldHWType
	fieldFlags
	fieldHWAddr
	fieldMask
	fieldDevice
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
	var line []ArpTableItem

	f, err := os.Open("/proc/net/arp")
	if err != nil {
		return line, fmt.Errorf("open file: %w", err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()

	for s.Scan() {
		fields := strings.Fields(s.Text())
		line = append(line, ArpTableItem{
			IPAddr: fields[fieldIPAddr],
			HWType: fields[fieldHWType],
			Flags:  fields[fieldFlags],
			HWAddr: fields[fieldHWAddr],
			Mask:   fields[fieldMask],
			Device: fields[fieldDevice],
		})
	}

	return line, nil
}
