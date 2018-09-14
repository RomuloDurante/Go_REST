package main

import (
	"net/http"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api"
	"github.com/RomuloDurante/WordHunter/restApi_GO/view"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}
	http.HandleFunc("/", view.HandleView)
	http.HandleFunc("/api/", api.HandleAPI)
	http.Handle("/css/", http.FileServer(http.Dir("view/public")))
	http.Handle("/img/", http.FileServer(http.Dir("view/public")))
	http.Handle("/script/", http.FileServer(http.Dir("view/public")))
	server.ListenAndServe()

}
