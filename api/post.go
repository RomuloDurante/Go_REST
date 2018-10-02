package api

import (
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/books"
)

// POST service
type POST struct{}

//BooksPOST service ...
func (p POST) BooksPOST(w http.ResponseWriter, r *http.Request) (err error) {
	//DataBook
	var dataBook *books.BookData
	//get the payload
	dataBook, newData, err := dataBook.Payload(r, "post")

	if err != nil {
		return err
	}

	// compare the data with newData and see if the author name matches
	author, opt := dataBook.CheckAuthor(map[string]string{"by": "name"}, &newData)

	// if the author exists append the newBook
	if opt == true {
		// check if book exits
		_, _, opt := dataBook.CheckBook(map[string]string{"by": "name"}, &newData)

		// if not create newBook
		if opt == false {
			dataBook.CreateBook(newData, author)
			w.Write([]byte("Book was create"))

		} else {
			w.Write([]byte("Book already exists"))
		}

		//if the author does not exists createAuthor
	} else if opt == false {
		dataBook.CreateAuthor(newData)
		w.Write([]byte("Author was create"))
	}

	return nil
}
