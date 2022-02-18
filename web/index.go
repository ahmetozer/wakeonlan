package web

import (
	"embed"
	"encoding/json"
	"net/http"
)

type page struct {
	file        string
	contentType string
}

var (
	//go:embed static
	static embed.FS
	pages  = map[string]page{
		"/":          {file: "static/index.html", contentType: "text/html"},
		"/style.css": {file: "static/style.css", contentType: "text/css"},
	}
)

// Index is the main entry point for the web server.
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	page, ok := pages[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	fileContent, err := static.ReadFile(page.file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	w.Header().Set("Content-Type", page.contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)
}
