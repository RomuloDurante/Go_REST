package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/books"
)

// GET service
type GET struct{}

/**************************************************/
/*         Deal with BookData                     */

// BooksGET ...
func (get GET) BooksGET(w http.ResponseWriter, r *http.Request, path map[string]string) (err error) {
	//DataBook
	var dataBook *books.BookData

	//get the payload
	dataBook, _, err = dataBook.Payload(r, "get")

	if err != nil {
		return err
	}

	// call the function and get the data
	var data []byte
	switch path["path"] {

	case "/api/dataBook/":
		data, err = get.GetAllBookData(dataBook)
		if err != nil {
			return err
		}
		break

	case "/api/dataBook/book/":
		//if query does not exists respose all books
		if path["id"] != "" {
			data, _, err = get.GetOneBook(dataBook, path["id"], "json")
			if err != nil {
				return err
			}
			break
		} else {
			data, _, err = get.GetAllBooks(dataBook, "json")
			if err != nil {
				return err
			}
			break
		}

	case "/api/dataBook/author/":
		// if query does not exists respose all author
		if path["id"] != "" {
			data, _, err = get.GetOneAuthor(dataBook, path["id"], "json")
			if err != nil {
				return err
			}
			break

		} else {
			data, _, err = get.GetAllAuthor(dataBook, "json")
			if err != nil {
				return err
			}
			break
		}
	case "/api/dataBook/rating/":
		if path["id"] != "" && path["rating"] != "" {

			err = dataBook.UpdateScore(path["id"], path["rating"])
			if err != nil {
				return err
			}
			w.Write([]byte("Book was Update"))
			break
		}
		w.Write([]byte("Could Not Update the Rating"))
		break
	default:
		return fmt.Errorf("%v", "Service does not match.")
	}
	// set content type as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.Write(data)

	return nil
}

//GetAllBookData -> domain/api/bookDataBook/
func (get GET) GetAllBookData(dataBook *books.BookData) ([]byte, error) {
	data, err := json.MarshalIndent(dataBook, "", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

//GetAllAuthor -> domain/api/dataBook/author/
func (get GET) GetAllAuthor(dataBook *books.BookData, op string) ([]byte, *[]books.Author, error) {

	author := dataBook.Author
	if op == "json" {
		data, err := json.MarshalIndent(author, "", "")
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		return data, nil, nil
	}

	return nil, &author, nil
}

//GetOneAuthor -> domain/api/dataBook/author/{id}
func (get GET) GetOneAuthor(dataBook *books.BookData, id string, op string) ([]byte, *books.Author, error) {

	author, opt := dataBook.CheckAuthor(map[string]string{"by": "query", "query": id}, nil)
	if opt == false {
		return nil, nil, fmt.Errorf("%v", "Could not find the Author")
	}

	if op == "json" {
		data, err := json.MarshalIndent(author, "", "")
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		return data, nil, err
	}

	return nil, author, nil
}

//GetAllBooks -> domain/api/dataBook/book/
func (get GET) GetAllBooks(dataBook *books.BookData, op string) ([]byte, *[]books.Book, error) {

	books := dataBook.Books

	if op == "json" {
		data, err := json.MarshalIndent(books, "", "")
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		return data, nil, nil
	}

	return nil, &books, nil
}

//GetOneBook -> domain/api/dataBook/book/{id}
func (get GET) GetOneBook(dataBook *books.BookData, id string, op string) ([]byte, *books.Book, error) {

	book, _, opt := dataBook.CheckBook(map[string]string{"by": "query", "query": id}, nil)

	if opt == false {
		return nil, nil, fmt.Errorf("%v", "Could not find the Book")
	}

	if op == "json" {
		data, err := json.MarshalIndent(book, "", "")
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		return data, nil, nil
	}

	return nil, book, nil
}

/*                   END BookData                     */
/******************************************************/
