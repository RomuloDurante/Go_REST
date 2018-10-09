package api

import (
	"net/http"

	"github.com/RomuloDurante/Go_REST/api/controller"
)

// HandleAPI func ...
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	var c *controller.Controller
	var err error

	//Controller start
	c = c.Start(r)

	item := c.Path["item"]

	if item == "user" || item == "book" || item == "author" {
		switch c.Method {
		case "GET":
			err = Get(w, r, c)
			break
		case "POST":
			err = Post(w, r, c)
			break
		case "PUT":
			err = Put(w, r, c)
			break
		case "DELETE":
			err = Delete(w, r, c)
			break
		}

	} else {
		w.WriteHeader(404)
		w.Write([]byte("Service does not matches"))

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
