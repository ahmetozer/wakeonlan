package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/ahmetozer/wakeonlan/share"
)

type wol struct {
	MAC  string
	IF   string
	ADDR string
	PORT string
}

type wolRespond struct {
	RequestNo int
	Status    string
}

func WakeOnLan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var PCs []wol

	err := json.NewDecoder(r.Body).Decode(&PCs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respond := []wolRespond{}

	for i, pc := range PCs {
		mac, err := net.ParseMAC(pc.MAC)
		if err != nil {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    err.Error(),
			})
			continue
		}
		err = share.MagicPacket{MAC: mac, IF: pc.IF, ADDR: pc.ADDR, PORT: pc.PORT}.SendMagicPacket()
		if err == nil {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    "packet send",
			})
		} else {
			respond = append(respond, wolRespond{
				RequestNo: i + 1,
				Status:    err.Error(),
			})
		}
	}

	jsonResp, err := json.Marshal(respond)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"status\":\"%v\",\"error\":\"%v\"}", http.StatusInternalServerError, err)))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
