package books

import (
	"net/http"
	"os"
)

// Delete service
func Delete(w http.ResponseWriter, r *http.Request) (err error) {
	query := r.URL.Query()
	var id string

	for key, value := range query {
		if key == "id" {
			for _, queryID := range value {
				id = queryID
			}
		}

	}

	// try use id to delete user
	err = os.Remove(".data/books/infoBook/" + id + ".json")
	if err != nil {
		return
	}

	err = os.Remove(".data/books/rawBook/" + id + ".txt")
	if err != nil {
		return
	}

	w.Write([]byte("book was deleted"))
	return nil
}
