package api

import (
	"fmt"
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/books"
)

// DELETE service
type DELETE struct{}

// BooksDELETE service
func (d DELETE) BooksDELETE(w http.ResponseWriter, r *http.Request, path map[string]string) (err error) {
	//DataBook
	var dataBook *books.BookData

	//get the payload
	dataBook, _, err = dataBook.Payload(r, "get")

	if err != nil {
		return err
	}

	// call the function and get the data
	switch path["path"] {
	case "/api/dataBook/book/":
		if path["id"] != "" {
			_, index, opt := dataBook.CheckBook(map[string]string{"by": "query", "query": path["id"]}, nil)
			if opt == false {
				return fmt.Errorf("%v", "Could not find the book")
			}

			dataBook.Books = append(dataBook.Books[:index], dataBook.Books[index+1:]...)
			fmt.Println(dataBook.Books) //TODO: test delete book
			w.Write([]byte("Book was deleted"))
			break
		}
		return fmt.Errorf("%v", "Could not find the book")
	}

	return nil
}
