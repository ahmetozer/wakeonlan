package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/ahmetozer/wakeonlan/share"
)

type wolRequest struct {
	HWAddr string
	Device string
	IPAddr string
	Port   string
}

type wolRespond struct {
	RequestNo int
	Status    string
}

// WakeOnLan is a handler for wake on lan.
func WakeOnLan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var wolTargets []wolRequest

	err := json.NewDecoder(r.Body).Decode(&wolTargets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	var respond []wolRespond

	for i, wolTarget := range wolTargets {
		mac, err := net.ParseMAC(wolTarget.HWAddr)
		if err != nil {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    fmt.Sprintf("invalid MAC address: %s", wolTarget.HWAddr),
			})
			continue
		}

		magicPacket := share.MagicPacket{
			HWAddr: mac,
			Device: wolTarget.Device,
			IPAddr: wolTarget.IPAddr,
			Port:   wolTarget.Port,
		}

		err = magicPacket.SendMagicPacket()
		if err == nil {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    "packet send",
			})
		} else {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    fmt.Sprintf("packet send error: %v", err),
			})
		}
	}

	err = json.NewEncoder(w).Encode(respond)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}
}
