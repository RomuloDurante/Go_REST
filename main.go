package main

import (
	"net/http"

	"github.com/RomuloDurante/Go_REST/api"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/api/", api.HandleAPI)
	server.ListenAndServe()

}
