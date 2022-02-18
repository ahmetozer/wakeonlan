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

const magicPacketSize = 102

// SendMagicPacket
// Send magic packet to destination to wake up the remote
func (mp *MagicPacket) SendMagicPacket() error {

	if len(mp.HWAddr) != 6 {
		return fmt.Errorf("invalid mac '%v'", mp.HWAddr)
	}

	var packet [magicPacketSize]byte
	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], mp.HWAddr)
		offset += 6
	}

	iface, err := net.InterfaceByName(mp.Device)
	if err != nil {
		return fmt.Errorf("interface by name: %w", err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return fmt.Errorf("interface addrs: %w", err)
	}

	dialer := &net.Dialer{
		LocalAddr: &net.UDPAddr{
			IP:   addrs[0].(*net.IPNet).IP,
			Port: 0,
		},
	}
	conn, err := dialer.Dial("udp", mp.IPAddr+":"+mp.Port)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write(packet[:])
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}
