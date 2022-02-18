package share

import (
	"net"
)

type Interface struct {
	Name   string
	IPAddr []string
}

// GetInterfaces
// Get list of the interfaces with or without IP addresses
func GetInterfaces() ([]Interface, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return []Interface{}, err
	}

	var netInterfaces []Interface

	for _, iface := range ifaces {

		addrs, err := iface.Addrs()
		if err != nil {
			netInterfaces = append(netInterfaces, Interface{
				Name: iface.Name,
			})
			continue
		}

		var ipAddrs []string

		for _, addr := range addrs {
			ipAddrs = append(ipAddrs, addr.String())
		}

		netInterfaces = append(netInterfaces, Interface{
			Name:   iface.Name,
			IPAddr: ipAddrs,
		})

	}

	return netInterfaces, nil
}
