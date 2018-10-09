package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	
	"github.com/RomuloDurante/Go_REST/api/helpers"
)

//Controller ...
type Controller struct {
	// Users
	DtUser *DataUser
	//Books
	DtBook *DataBook
	//resources
	Method string
	Query  map[string][]string
	Path   map[string]string // service | model | id

	//helpers
	helpers.Helpers
}

// Start Controller
func (c *Controller) Start(r *http.Request) *Controller {

	p := make([]string, 10)

	splitPath := strings.Split(r.URL.Path, "/")
	for key, value := range splitPath {
		p[key] = value
	}

	return &Controller{
		Method: r.Method,
		Query:  r.URL.Query(),
		Path: map[string]string{
			"method": r.Method,
			"item":   p[2],
			"id":     p[3],
			"rating": p[4],
		},
	}
}

//ReadData ...
func (c *Controller) ReadData() (data []byte, err error) {
	if c.Path["item"] == "book" || c.Path["item"] == "author" {
		data, err = ioutil.ReadFile(".data/dataBook/dataBook.json")

	} else if c.Path["item"] == "user" {
		data, err = ioutil.ReadFile(".data/dataUser/dataUser.json")
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

//GetBody ...
func (c *Controller) GetBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return body, nil
}

//CreateData ..
func (c *Controller) CreateData(data []byte) error {
	var err error

	if c.Path["item"] == "book" || c.Path["item"] == "author" {
		err = ioutil.WriteFile(".data/dataBook/dataBook.json", data, 0600)

	} else if c.Path["item"] == "user" {
		err = ioutil.WriteFile(".data/dataUser/dataUser.json", data, 0600)
	}

	if err != nil {
		return err
	}

	return nil
}

//GetItem entire item and by ID...
func (c *Controller) GetItem() ([]byte, error) {

	data, err := c.ReadData()

	if err != nil {
		return nil, err
	}

	switch c.Path["item"] {
	case "book":
		//unmarshal the data
		err := json.Unmarshal(data, &c.DtBook)
		if err != nil {
			return nil, err
		}

		if c.Path["id"] == "" {
			//return the entire Item
			dt, err := c.JSONstring(c.DtBook.TakeBooks())
			if err != nil {
				return nil, err
			}
			return dt, nil
		}
		//otherwise return the item by ID
		dt, err := c.JSONstring(c.DtBook.ByID("book", c.Path["id"]))
		if err != nil {
			return nil, err
		}
		return dt, nil

	case "author":
		err := json.Unmarshal(data, &c.DtBook)
		if err != nil {
			return nil, err
		}

		if c.Path["id"] == "" {
			//return the entire Item
			dt, err := c.JSONstring(c.DtBook.TakeAuthor())
			if err != nil {
				return nil, err
			}
			return dt, nil
		}
		//otherwise return the item by ID
		dt, err := c.JSONstring(c.DtBook.ByID("author", c.Path["id"]))
		if err != nil {
			return nil, err
		}
		return dt, nil
	case "user":
		//umarshal the data
		err := json.Unmarshal(data, &c.DtUser)
		if err != nil {
			return nil, err
		}

		if c.Path["id"] == "" {
			//return the entire Item
			dt, err := c.JSONstring(c.DtUser.TakeUser())
			if err != nil {
				return nil, err
			}
			return dt, nil
		}
		//otherwise return the item by ID
		dt, err := c.JSONstring(c.DtUser.ByID("user", c.Path["id"]))
		if err != nil {
			return nil, err
		}
		return dt, nil

	}

	return nil, nil
}

//CreateItem ...
func (c *Controller) CreateItem(old, new []byte) ([]byte, error) {
	if c.Path["item"] == "book" {
		var newBook DataBook
		err := json.Unmarshal(old, &c.DtBook)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(new, &newBook)
		if err != nil {
			return nil, err
		}

		authorID, err := c.DtBook.ByName("author", newBook.Author[0].Name)
		if err != nil {
			return nil, fmt.Errorf("The author does not exists") //TODO: we need create the author redirect to -> http//.../author
		}

		_, err = c.DtBook.ByName("book", newBook.Books[0].Name)
		if err != nil { //book does not exists -> create the book
			book := newBook.Books[0]
			book.ID = c.Token(10)
			book.AuthorID = authorID
			c.DtBook.Books = append(c.DtBook.Books, book)

			dt, err := c.JSONstring(c.DtBook, nil)
			if err != nil {
				return nil, err
			}

			err = c.CreateData(dt)
			if err != nil {
				return nil, err
			}

			return []byte("Book was create"), nil
		}
		return nil, fmt.Errorf("The Book already exists")

	} else if c.Path["item"] == "author" {
		var newBook DataBook
		err := json.Unmarshal(old, &c.DtBook)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(new, &newBook)
		if err != nil {
			return nil, err
		}

		_, err = c.DtBook.ByName("author", newBook.Author[0].Name)
		if err != nil {
			author := newBook.Author[0]
			author.ID = c.Token(10)
			c.DtBook.Author = append(c.DtBook.Author, author)

			dt, err := c.JSONstring(c.DtBook, nil)
			if err != nil {
				return nil, err
			}

			err = c.CreateData(dt)
			if err != nil {
				return nil, err
			}
			return []byte("Author was create"), nil
		}

		return nil, fmt.Errorf("The Author already exists")
	}

	return nil, nil
}
