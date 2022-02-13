package share

import (
	"fmt"
	"net"
)

type MagicPacket struct {
	MAC  net.HardwareAddr
	IF   string
	ADDR string
	PORT string
}

// SendMagicPacket
// Send magic packet to destination to wake up the remote
func (MP MagicPacket) SendMagicPacket() error {

	if len(MP.MAC) != 6 {
		return fmt.Errorf("invalid mac '%v'", MP.MAC)
	}

	var packet [102]byte
	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], MP.MAC)
		offset += 6
	}

	ief, err := net.InterfaceByName(MP.IF)
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
	conn, err := dialer.Dial("udp", MP.ADDR+":"+MP.PORT)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet[:])
	return err
}
