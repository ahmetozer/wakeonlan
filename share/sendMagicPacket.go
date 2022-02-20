package share

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type MagicPacket struct {
	HWAddr net.HardwareAddr
	Device string
	IPAddr string
	Port   string
}

const magicPacketSize = 102

// SendMagicPacket
// Send magic packet to remote to wake up the remote
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

	// To determine remote Address Is IPv6
	remoteIsIPv6 := strings.Contains(mp.IPAddr, ":")
	remote := mp.IPAddr
	sourceIP := net.IP{}

	// If remote address is IPv6, find same net type IP address
	if remoteIsIPv6 {
		sourceIP, err = findIP6SourceAddress(&addrs, mp.IPAddr)

		// no ip v6 addr is found
		if err != nil {

			return fmt.Errorf("find IP6 source address: interface %v: %w", mp.Device, err)
		}
		if sourceIP.IsLinkLocalUnicast() {
			remote = fmt.Sprintf("[%v%%%v]", remote, mp.Device)
		} else {
			remote = fmt.Sprintf("[%v]", remote)
		}

	} else {
		IPv4Found := false
		for i := range addrs {
			if IP := addrs[i].(*net.IPNet).IP; strings.Contains(IP.String(), ".") {
				IPv4Found = true
				sourceIP = IP
				break
			}
		}
		if !IPv4Found {

			return fmt.Errorf("interface '%v' does not have a IPv4 addr", mp.Device)
		}
	}

	sourceAddr := &net.UDPAddr{
		IP:   sourceIP,
		Port: 0, // Port zero is auto select for golang
		Zone: mp.Device,
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", remote+":"+mp.Port)
	if err != nil {

		return fmt.Errorf("remote resolve error: %w", err)
	}

	conn, err := net.DialUDP("udp", sourceAddr, remoteAddr)

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

func findIP6SourceAddress(addrs *[]net.Addr, remoteAddress string) (net.IP, error) {

	dstIP := net.ParseIP(remoteAddress)
	// Try to find same net type IP
	if dstIP.IsLinkLocalUnicast() { // if it is link local
		for i := range *addrs {
			if IP := (*addrs)[i].(*net.IPNet).IP; IP.IsLinkLocalUnicast() && strings.Contains(IP.String(), ":") {
				return IP, nil
			}
		}
		return net.IP{}, errors.New("no link local unicast IPv6 address is found")

	}
	for i := range *addrs {
		if IP := (*addrs)[i].(*net.IPNet).IP; IP.IsGlobalUnicast() && strings.Contains(IP.String(), ":") {
			return IP, nil
		}
	}
	return net.IP{}, errors.New("no global IPv6 address is found")
}
