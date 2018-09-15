package api

import (
	"net/http"
	"strings"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api/books"
	"github.com/RomuloDurante/WordHunter/restApi_GO/api/users"
)

// HandleAPI func ...
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	var err error

	path := strings.Split(r.URL.Path, "/")

	if path[2] == "user" {
		switch r.Method {
		case "GET":
			err = users.Get(w, r)
		case "POST":
			err = users.Post(w, r)
		case "PUT":
			err = users.Put(w, r)
		case "DELETE":
			err = users.Delete(w, r)
		}
	} else if path[2] == "book" {
		switch r.Method {
		case "GET":
			err = books.Get(w, r)
		case "POST":
			err = books.Post(w, r)
		case "PUT":
			err = books.Put(w, r)
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
