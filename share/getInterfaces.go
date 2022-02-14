package share

import (
	"net"
)

type Interface struct {
	Device string
	IPAddr []string
}

// GetInterfaces
// Get list of the interfaces with IP addresses
func GetInterfaces() ([]Interface, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return []Interface{}, err
	}
	tempInterfaces := []Interface{}

	for _, i := range ifaces {

		addrs, err := i.Addrs()

		if err != nil {
			tempInterfaces = append(tempInterfaces, Interface{
				Device: i.Name,
			})
			continue
		}
		tempIfaddrs := []string{}

		for _, a := range addrs {
			tempIfaddrs = append(tempIfaddrs, a.String())
		}

		tempInterfaces = append(tempInterfaces, Interface{
			Device: i.Name,
			IPAddr: tempIfaddrs,
		})

	}

	return tempInterfaces, nil
}
