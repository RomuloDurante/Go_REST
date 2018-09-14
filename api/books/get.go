package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get service
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	query := r.URL.Query()
	var id string

	for key, value := range query {
		if key == "id" {
			for _, queryID := range value {
				id = queryID
			}
		}

	}

	infoBook, err := ioutil.ReadFile(".data/books/infoBook/" + id + ".json")
	if err != nil {
		return
	}
	var book Book
	json.Unmarshal(infoBook, &book)

	rawBook, err := ioutil.ReadFile(".data/books/rawBook/" + id + ".txt")
	if err != nil {
		return
	}
	bookSlice := strings.Split(string(rawBook), "")

	fmt.Println(bookSlice[2])
	fmt.Println(book)
	w.Header().Set("Content-type", "text/plain")
	return nil
}
