package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmetozer/wakeonlan/handlers"
	"github.com/ahmetozer/wakeonlan/web"
)

var LISTEN = ""

func init() {
	LISTEN = os.Getenv("LISTEN")
	if LISTEN == "" {
		LISTEN = ":8080"
	}
}

func main() {

	http.HandleFunc("/", web.Index)

	http.HandleFunc("/api/interfaces", handlers.Interfaces)

	http.HandleFunc("/api/arpentries", handlers.ArpEntries)

	http.HandleFunc("/api/wakeonlan", handlers.WakeOnLan)

	log.Printf("Starting http server on %q \n", LISTEN)
	err := http.ListenAndServe(LISTEN, nil)
	if err != nil {
		log.Fatal(err)
	}
}
