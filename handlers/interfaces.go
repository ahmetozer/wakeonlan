package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmetozer/wakeonlan/share"
)

// Interfaces
// Serve interface with details
func Interfaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	interfaces, err := share.GetInterfaces()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	err = json.NewEncoder(w).Encode(interfaces)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}
}
