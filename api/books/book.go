package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/RomuloDurante/WordHunter/api/helpers"
)

// BookData ... -> MAIN DATA
type BookData struct {
	BookInfo infoBook `json:"bookInfo"`
	Author   []Author `json:"author"`
	Books    []Book   `json:"books"`
}

// infobook
type infoBook struct {
	NumberOfBooks   int      `json:"numberOfBooks"`
	NumberOfAuthors int      `json:"numberOfAuthors"`
	FullRating      []string `json:"fullRating"`
}

//Author ...
type Author struct {
	Name        string  `json:"name"`
	ID          string  `json:"id"`
	QueryBook   string  `json:"queryBook,omitempty"`
	AboutAuthor string  `json:"aboutAuthor"`
	Forum       []Forum `json:"forum"`
}

// Book ...
type Book struct {
	Name      string  `json:"name"`
	ID        string  `json:"id"`
	AuthorID  string  `json:"authorId"`
	Rating    int     `json:"rating"`
	Year      string  `json:"year"`
	AboutBook string  `json:"aboutBook"`
	Dificulty string  `json:"dificulty"`
	Forum     []Forum `json:"forum"`
}

//Forum ...
type Forum struct { //TODO: forum

}

// END Author

/*************Methods BOOKDATA*****************************/

//Payload ...
func (dataBook *BookData) Payload(r *http.Request, opt string) (*BookData, BookData, error) {
	//open the data books
	books, err := ioutil.ReadFile(".data/books/infoBook/book.json")
	json.Unmarshal(books, &dataBook)

	// create newdata (*BookData, BookData, error)
	var newData BookData

	if opt == "post" {
		// read the body content
		body := helpers.GetBody(r)

		// push the body into newData
		err = json.Unmarshal(body, &newData)
		if err != nil {
			return nil, newData, err
		}

		return dataBook, newData, nil

	}

	return dataBook, newData, nil
}

// CheckRating ..
func (dataBook *BookData) CheckRating() []Book {

	var rating []Book

	//range in the rating and push the books on the view rating
	for key, id := range dataBook.BookInfo.FullRating {
		// find the book
		book, _, _ := dataBook.CheckBook(map[string]string{"by": "query", "query": id}, nil)

		rating = append(rating, *book)

		fmt.Println(rating[key].Rating) //TODO: return the rating

	}

	return rating

}

/***********************About Author *************************/

// CheckAuthor check if author exists
func (dataBook *BookData) CheckAuthor(opt map[string]string, newData *BookData) (*Author, bool) {

	if opt["by"] == "name" {
		for _, author := range dataBook.Author {
			for _, newAuthor := range newData.Author {
				if author.Name == newAuthor.Name {
					return &author, true
				}
			}
		}
	} else if opt["by"] == "query" {
		for _, author := range dataBook.Author {
			if author.ID == opt["query"] {
				return &author, true
			}
		}
	}

	return nil, false
}

// CreateAuthor ...
func (dataBook *BookData) CreateAuthor(newData BookData) error {
	for _, author := range newData.Author {
		author.ID = helpers.Token(10)

		dataBook.Author = append(dataBook.Author, author)
		dataBook.BookInfo.NumberOfAuthors++
	}
	err := dataBook.Create()

	if err != nil {
		return err
	}

	return nil
}

//***->End author

/***********************About Book *************************/

// CheckBook check if author exists
func (dataBook *BookData) CheckBook(opt map[string]string, newData *BookData) (*Book, int, bool) {

	if opt["by"] == "name" {
		for _, book := range dataBook.Books {
			for _, newBook := range newData.Books {
				if book.Name == newBook.Name {
					return &book, 0, true
				}
			}
		}
	} else if opt["by"] == "query" {
		for key, book := range dataBook.Books {
			if book.ID == opt["query"] {
				return &book, key, true
			}
		}
	}

	return nil, 0, false
}

// CreateBook ...
func (dataBook *BookData) CreateBook(newData BookData, author *Author) error {
	for _, book := range newData.Books {
		book.ID = helpers.Token(10)

		book.AuthorID = author.ID
		dataBook.Books = append(dataBook.Books, book)
	}

	err := dataBook.Create()

	if err != nil {
		return err
	}

	return nil
}

//GetAllBooks -> domain/api/dataBook/book/
func (dataBook *BookData) GetAllBooks(op string) ([]byte, *[]Book, error) {

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
func (dataBook *BookData) GetOneBook(id string, op string) ([]byte, *Book, error) {

	book, _, opt := dataBook.CheckBook(map[string]string{"by": "query", "query": id}, nil)

	if opt == false {
		return nil, nil, fmt.Errorf("%v", "Could not find the Author")
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

//***->End Book

/*********************** -> OTHERS METHODS
updateScore ->// domain/api/bookData/rating/book/{id}/{rating}{in this format n1-n2: n1[0 "minus", 1 "plus"] n2[number n to sum o subtract of the rating]} TODO:update score
EX -> localhost:5000/api/bookData/rating/yL5uU8wW8kO7uM5yC2oU5uH5iL8oE5/1-2 (increases the value of n2 to book rating)
***/

// UpdateScore update the rating and fullRating
func (dataBook *BookData) UpdateScore(id string, r string) error {

	rating := strings.Split(r, "-")
	var operator int
	var note int

	// deal with possible errors
	if len(rating) > 1 && rating[1] != "" {
		operator, _ = strconv.Atoi(rating[0]) // 0 -> subtration 1 -> sum
		note, _ = strconv.Atoi(rating[1])
	} else {
		return fmt.Errorf("%v", "Could not update the book rating")
	}

	book, index, opt := dataBook.CheckBook(map[string]string{"by": "query", "query": id}, nil)

	if opt == false {
		return fmt.Errorf("%v", "Could not update the book rating")
	}

	// if operator = 0 minus otherwise plus
	if operator == 0 {
		book.Rating -= note
	} else {
		book.Rating += note
	}
	//TODO: find way to organize the fullrating with the new vote
	//push the book with the new rating into databook
	dataBook.Books[index] = *book

	err := dataBook.Create()

	if err != nil {
		return err
	}

	return nil
}

/***********************Create Data *************************/

// Create ...
func (dataBook *BookData) Create() error {
	// Marshal newDataBook
	data, err := json.MarshalIndent(dataBook, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}
	//update the dataBook
	err = ioutil.WriteFile(".data/books/infoBook/book.json", data, 0666)
	if err != nil {
		return err
	}

	return nil
}

//***->End Data
