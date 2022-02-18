package handlers

import (
	"encoding/json"
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

	arpTable, err := share.GetArpTable()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	err = json.NewEncoder(w).Encode(arpTable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}
}
