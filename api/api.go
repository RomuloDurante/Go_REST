package api

import (
	"fmt"
	"net/http"
	"strings"
)

// API ...
type API struct {
	//Services
	GET
	POST
	PUT
	DELETE
	//resources
	method  string
	service string
	query   map[string][]string
	path    map[string]string // service | obj | id
}

// Start the API services
func (api *API) Start(r *http.Request) *API {

	p := make([]string, 10)

	splitPath := strings.Split(r.URL.Path, "/")
	for key, value := range splitPath {
		p[key] = value
	}

	var switchPath = func(p []string) string {
		path := "/api/" + p[2] + "/"
		if p[3] != "" {
			path += p[3] + "/"
		}
		return path
	}

	return &API{
		method:  r.Method,
		service: strings.Split(r.URL.Path, "/")[2],
		query:   r.URL.Query(),
		path: map[string]string{
			"service": p[2],
			"obj":     p[3],
			"id":      p[4],
			"rating":  p[5],
			"path":    switchPath(p), // Create the path to choice the methods
		},
	}
}

// HandleAPI func ...
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	var api *API
	var err error

	//start API
	api = api.Start(r)

	if api.path["service"] == "userData" {
		switch api.method {
		case "GET":
			// users.Get(w, r)
			break
		case "POST":
			fmt.Println("post")
			break
		case "PUT":
			fmt.Println("put")
			break
		case "DELETE":
			fmt.Println("delete")
			break
		}

	} else if api.path["service"] == "dataBook" {
		switch api.method {
		case "GET":
			err = api.BooksGET(w, r, api.path)
			break
		case "POST":
			err = api.BooksPOST(w, r)
			break
		case "PUT":
			err = api.BooksPUT(w, r, api.path)
			break
		case "DELETE":
			err = api.BooksDELETE(w, r, api.path)
			break
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
