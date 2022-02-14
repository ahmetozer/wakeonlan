package share

import (
	"fmt"
	"net"
)

type MagicPacket struct {
	HWAddr net.HardwareAddr
	Device string
	IPAddr string
	Port   string
}

// SendMagicPacket
// Send magic packet to destination to wake up the remote
func (MP MagicPacket) SendMagicPacket() error {

	if len(MP.HWAddr) != 6 {
		return fmt.Errorf("invalid mac '%v'", MP.HWAddr)
	}

	var packet [102]byte
	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], MP.HWAddr)
		offset += 6
	}

	ief, err := net.InterfaceByName(MP.Device)
	if err != nil {
		return err
	}
	addrs, err := ief.Addrs()
	if err != nil {
		return err
	}

	dialer := &net.Dialer{
		LocalAddr: &net.UDPAddr{
			IP:   addrs[0].(*net.IPNet).IP,
			Port: 500,
		},
	}
	conn, err := dialer.Dial("udp", MP.IPAddr+":"+MP.Port)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet[:])
	return err
}
