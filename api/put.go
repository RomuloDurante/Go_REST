package api

import (
	"net/http"

	"github.com/RomuloDurante/Go_REST/api/controller"
)

//Put ...
func Put(w http.ResponseWriter, r *http.Request, c *controller.Controller) error {

	data, _ := c.GetItem()

	w.Write(data)
	return nil
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/RomuloDurante/WordHunter/api/model"
// )

// // PUT service
// type PUT struct {
// 	GET // use the methods from GET service to take data
// }

// /**************************************************/
// /*         Deal with DataBook                     */

// // BooksPUT service
// func (put PUT) BooksPUT(w http.ResponseWriter, r *http.Request, path map[string]string) (err error) {
// 	//DataBook
// 	var dataBook *model.DataBook

// 	//get the payload
// 	_, newData, err := dataBook.Payload(r, "post")

// 	// get the data from database
// 	dataBook, _, err = dataBook.Payload(r, "get")

// 	if err != nil {
// 		return err
// 	}

// 	switch path["obj"] {
// 	case "book":
// 		err = put.UpdateBook(dataBook, newData, path)
// 		break
// 	case "author":
// 		err = put.UpdateAuthor(dataBook, newData, path)
// 		break
// 	}

// 	if err != nil {
// 		return err

// 	}
// 	w.Write([]byte("Data was updated"))

// 	return nil
// }

// // UpdateBook ...
// func (put PUT) UpdateBook(dataBook *model.DataBook, newData model.DataBook, path map[string]string) error {
// 	var opt bool

// 	// find the index of the book
// 	_, index, opt := dataBook.CheckBook(map[string]string{"by": "query", "query": path["id"]}, &newData)
// 	if opt == false {
// 		return fmt.Errorf("%v", "Could not find the Book")
// 	}

// 	//return book from dataBase
// 	byteBook, _, _ := put.GetOneBook(dataBook, path["id"], "json") // we use json and type []byte to compare the structs

// 	// get the newBook from request
// 	newByteBook, newBook, _ := put.GetOneBook(&newData, path["id"], "json") // we use json and type []byte to compare the structs

// 	//compare the book and look if the data matches
// 	for i := 0; i < len(byteBook); i++ {
// 		if byteBook[i] != newByteBook[i] {
// 			opt = false
// 			break
// 		} else {
// 			opt = true
// 		}
// 	}

// 	// unmarshal the byteBook into Book
// 	if opt == false {
// 		err := json.Unmarshal(newByteBook, &newBook)
// 		if err != nil {
// 			return err
// 		}

// 		//insert the newBook on the dataBook
// 		dataBook.Books[index] = *newBook
// 		//update hte data
// 		err = dataBook.Create()

// 		if err != nil {
// 			return err
// 		}
// 	} else if opt == true {
// 		return fmt.Errorf("%v", "Book is already update")
// 	}

// 	return nil

// }

// // UpdateAuthor ...
// func (put PUT) UpdateAuthor(dataBook *model.DataBook, newData model.DataBook, path map[string]string) error {
// 	var opt bool

// 	// find the index of the book
// 	_, index, opt := dataBook.CheckAuthor(map[string]string{"by": "query", "query": path["id"]}, &newData)

// 	if opt == false {
// 		return fmt.Errorf("%v", "Could not find the Author")
// 	}

// 	//return author from dataBase
// 	byteAuthor, _, _ := put.GetOneAuthor(dataBook, path["id"], "json") // we use json and type []byte to compare the structs

// 	// get the newAuthor from request
// 	newByteAuthor, newAuthor, _ := put.GetOneAuthor(&newData, path["id"], "json") // we use json and type []byte to compare the structs

// 	//compare the book and look if the data matches
// 	for i := 0; i < len(byteAuthor); i++ {
// 		if byteAuthor[i] != newByteAuthor[i] {
// 			opt = false
// 			break
// 		} else {
// 			opt = true
// 		}
// 	}

// 	// unmarshal the byteBook into Book
// 	if opt == false {
// 		err := json.Unmarshal(newByteAuthor, &newAuthor)
// 		if err != nil {
// 			return err
// 		}

// 		//insert the newBook on the dataBook
// 		dataBook.Author[index] = *newAuthor
// 		//update hte data
// 		err = dataBook.Create()

// 		if err != nil {
// 			return err
// 		}
// 	} else if opt == true {
// 		return fmt.Errorf("%v", "Author is already update")
// 	}

// 	return nil
// }
