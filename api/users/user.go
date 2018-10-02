package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/helpers"
)

// UserData ...
type UserData struct {
	UserInfo infoUsers `json:"userInfo"`
	Users    []User    `json:"users"`
}

// InfoUsers
type infoUsers struct {
	LastID        int `json:"lastId"` // TODO: I don't need the last ID because i will use the toke functions from helpers
	NumberOfUsers int `json:"numberOfUsers"`
	ActiveUsers   int `json:"activeUsers"`
	DetetedUsers  int `json:"detetedUsers "`
}

/***********************************/

// User struct
type User struct {
	ActiveUser bool              `json:"activeUser"`
	ID         string            `json:"id"`
	FirstName  string            `json:"firstName"`
	LastName   string            `json:"lastName"`
	Email      string            `json:"email"`
	Password   string            `json:"password"`
	Agree      bool              `json:"agree"`
	WordVault  map[string]string `json:"wordVault"`
	MyBooks    []myBooks         `json:"myBooks"`
}

//myBooks
type myBooks struct {
	ID       int `json:"id"`
	LastPage int `json:"lastPage"`
	//TODO: notes ?
}

/*************Methods USERDATA*****************************/

//Payload ...
func (userData *UserData) Payload(r *http.Request, opt string) (*UserData, UserData, error) {
	//open the data books
	books, err := ioutil.ReadFile(".data/users/users.json")
	json.Unmarshal(books, &userData)

	// create newdata (*UserData, UserData, error)
	var newData UserData

	if opt == "post" {
		// read the body content
		body := helpers.GetBody(r)

		// push the body into newData
		err = json.Unmarshal(body, &newData)
		if err != nil {
			return nil, newData, err
		}

		return userData, newData, nil

	}

	return userData, newData, nil
}
