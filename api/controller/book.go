package controller

import (
	"fmt"
)

// DataBook ... -> MAIN DATA
type DataBook struct {
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

/*************Methods DataBook*****************************/

//ShowItem ...
func (dataBook *DataBook) ShowItem() {
	fmt.Printf("%+v\n", dataBook)
}

// TakeBooks ...
func (dataBook *DataBook) TakeBooks() ([]Book, error) {
	return dataBook.Books, nil
}

//TakeAuthor ...
func (dataBook *DataBook) TakeAuthor() ([]Author, error) {
	return dataBook.Author, nil
}

//ByID ...
func (dataBook *DataBook) ByID(item, id string) (interface{}, error) {
	if item == "book" {
		for _, book := range dataBook.Books {
			if book.ID == id {
				return &book, nil
			}
		}
	} else if item == "author" {
		for _, author := range dataBook.Author {
			if author.ID == id {
				return &author, nil
			}
		}
	}

	return nil, fmt.Errorf("%v", "Could not find the data by ID")
}

//ByName ...
func (dataBook *DataBook) ByName(item, name string) (string, error) {
	if item == "book" {
		for _, book := range dataBook.Books {
			if book.Name == name {
				return book.ID, nil
			}
		}
	} else if item == "author" {
		for _, author := range dataBook.Author {
			if author.Name == name {
				return author.ID, nil
			}
		}
	}

	return "", fmt.Errorf("%v", "Could not find the ItemByName")
}
