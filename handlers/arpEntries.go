package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahmetozer/wakeonlan/share"
)

// ArpEntries
// Serve mac addresses
func ArpEntries(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ArpTable, err := share.GetArpTable()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"status\":\"%v\",\"error\":\"%v\"}", http.StatusInternalServerError, err)))
		return
	}

	jsonResp, err := json.Marshal(ArpTable)
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
