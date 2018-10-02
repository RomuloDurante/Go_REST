package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/books"
)

// PUT service
type PUT struct {
	GET // use the methods from GET service to take data
}

/**************************************************/
/*         Deal with BookData                     */

// BooksPUT service
func (put PUT) BooksPUT(w http.ResponseWriter, r *http.Request, path map[string]string) (err error) {
	//DataBook
	var dataBook *books.BookData

	//get the payload
	_, newData, err := dataBook.Payload(r, "post")

	// get the data from database
	dataBook, _, err = dataBook.Payload(r, "get")

	if err != nil {
		return err
	}

	switch path["obj"] {
	case "book":
		err = put.UpdateBook(dataBook, newData, path)
		break
	case "author":
		err = put.UpdateBook(dataBook, newData, path)
		break
	}

	if err != nil {
		w.Write([]byte("Cold not Update the book"))

	} else {
		w.Write([]byte("Book was updated"))
	}

	return nil
}

// UpdateBook ...
func (put PUT) UpdateBook(dataBook *books.BookData, newData books.BookData, path map[string]string) error {
	var opt bool

	// find the index of the book
	_, index, _ := dataBook.CheckBook(map[string]string{"by": "query", "query": path["id"]}, &newData)

	//return book from dataBase
	byteBook, _, _ := put.GetOneBook(dataBook, path["id"], "json") // we use json and type []byte to compare the structs

	// get the newBook from request
	newByteBook, newBook, _ := put.GetOneBook(&newData, path["id"], "json") // we use json and type []byte to compare the structs

	//compare the book and look if the data matches
	for i := 0; i < len(byteBook); i++ {
		if byteBook[i] != newByteBook[i] {
			opt = false
			break
		} else {
			opt = true
		}
	}

	// unmarshal the byteBook into Book
	if opt == false {
		err := json.Unmarshal(newByteBook, &newBook)
		if err != nil {
			return err
		}

		//insert the newBook on the dataBook
		dataBook.Books[index] = *newBook
		//update hte data
		err = dataBook.Create()

		if err != nil {
			return err
		}
	} else if opt == true {
		return fmt.Errorf("%v", "")
	}

	return nil

}

// UpdateAuthor ...
func (put PUT) UpdateAuthor(dataBook *books.BookData) error {
	data, err := put.GetAllBookData(dataBook)

	fmt.Println(data, err)
	return nil
}
