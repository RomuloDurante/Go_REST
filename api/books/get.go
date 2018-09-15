package books

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Get service
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	// open the data books
	books, err := ioutil.ReadFile(".data/books/infoBook/books.json")

	if err != nil {
		return
	}
	// create and unmarshal the dataBook
	var checkBook BookData
	json.Unmarshal(books, &checkBook)

	// get the id from url
	query := r.URL.Query()
	var id string
	for key, value := range query {
		if key == "id" {
			for _, queryID := range value {
				id = queryID
			}
		}
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	// find the use that will be show to us
	var showBook Book
	for _, book := range checkBook.Books {
		if intID == book.ID {
			showBook = book

			// Json showBook
			data, err := json.MarshalIndent(showBook, "", "")
			if err != nil {
				return err
			}

			w.Header().Set("Content-type", "aplication/json")
			w.Write(data)

		}

	}
	return nil

	// infoBook, err := ioutil.ReadFile(".data/books/infoBook/" + id + ".json")
	// if err != nil {
	// 	return
	// }
	// var book Book
	// json.Unmarshal(infoBook, &book)

	// rawBook, err := ioutil.ReadFile(".data/books/rawBook/" + id + ".txt")
	// if err != nil {
	// 	return
	// }
	// bookSlice := strings.Split(string(rawBook), "")

	// fmt.Println(bookSlice[2])
	// fmt.Println(book)
	// w.Header().Set("Content-type", "text/plain")
	// return nil
}
