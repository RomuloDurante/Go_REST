package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api/helpers"
)

// Post service
func Post(w http.ResponseWriter, r *http.Request) (err error) {
	// open the data books
	books, err := ioutil.ReadFile(".data/books/infoBook/books.json")

	if err != nil {
		return
	}
	// create and unmarshal the dataBook
	var dataBook BookData
	json.Unmarshal(books, &dataBook)

	// create new book
	var newBook Book

	// create the new id
	id := dataBook.BookInfo.LastID + 1
	// push the new id to bookInfo
	dataBook.BookInfo.LastID = id

	// use the new id to create newUBook
	newBook.ID = id

	// read the body content
	body := helpers.GetBody(r)

	// push the body into newBook
	err = json.Unmarshal(body, &newBook)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// push newBook to dataBook
	dataBook.Books = append(dataBook.Books, newBook)
	dataBook.BookInfo.NumberOfBooks++

	// parse the dataBook to json
	data, err := json.MarshalIndent(dataBook, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// update the book data
	err = ioutil.WriteFile(".data/books/infoBook/books.json", data, 0666)
	if err != nil {
		return
	}

	w.Write([]byte("Book was create"))
	return nil
}
