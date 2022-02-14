package web

import (
	"embed"
	"fmt"
	"net/http"
)

type page struct {
	file  string
	ctype string
}

var (
	//go:embed static
	static embed.FS
	pages  = map[string]page{
		"/":          {file: "static/index.html", ctype: "text/html"},
		"/style.css": {file: "static/style.css", ctype: "text/css"},
	}
)

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

	p, err := static.ReadFile(page.file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"status\":\"%v\",\"error\":\"%v\"}", http.StatusInternalServerError, err)))
		return
	}
	w.Header().Set("Content-Type", page.ctype)
	w.WriteHeader(http.StatusOK)
	w.Write(p)
}
