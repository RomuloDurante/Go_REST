package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api/helpers"
)

// Post service
func Post(w http.ResponseWriter, r *http.Request) (err error) {
	var newBook Book

	// get the config file
	config := helpers.CheckConfig()
	// create the new id
	id := config.Book.ID + 1
	// push the new id to config file
	config.Book.ID = id
	// update the config file with the new id
	helpers.UpdateConfig(config)

	// use the new id to create newBook
	newBook.ID = id

	// read the body content
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	// push the body into newUser
	err = json.Unmarshal(body, &newBook)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// parse the newUser to json
	data, err := json.MarshalIndent(newBook, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// use the id to create new user
	bookID := strconv.Itoa(id)
	ioutil.WriteFile(".data/books/infoBook/"+bookID+".json", data, 0666)
	// os.Rename(userID+".json", ".data/users/"+userID+".json")

	w.Write([]byte("Book was create"))
	return nil
}
