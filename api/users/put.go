package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api/helpers"
)

// Put service
func Put(w http.ResponseWriter, r *http.Request) (err error) {
	// open the user file
	users, err := ioutil.ReadFile(".data/users/users.json")
	if err != nil {
		return
	}
	var checkUser UserData
	json.Unmarshal(users, &checkUser)

	// get the query string
	query := r.URL.Query()
	var id string

	for key, value := range query {
		if key == "id" {
			for _, queryID := range value {
				id = queryID
			}
		}
	}
	// conv id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	// get the body from payload
	body := helpers.GetBody(r)
	// user to update
	var updatedUser User
	// push the body into newUser
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// verify if the id exists
	for key, value := range checkUser.Users {
		if intID == value.ID && value.ActiveUser == true {

			// check if the data send to server matches with data is already there, if not update the data
			for {
				if checkUser.Users[key].FirstName != updatedUser.FirstName {
					checkUser.Users[key].FirstName = updatedUser.FirstName
					continue
				} else if checkUser.Users[key].LastName != updatedUser.LastName {
					checkUser.Users[key].LastName = updatedUser.LastName
					continue
				} else if checkUser.Users[key].Email != updatedUser.Email {
					checkUser.Users[key].Email = updatedUser.Email
					continue
				} else if checkUser.Users[key].Password != updatedUser.Password {
					checkUser.Users[key].Password = updatedUser.Password
				}
				break
			}
		}

	}

	// create update user
	upUser, err := json.MarshalIndent(checkUser, "", "")
	if err != nil {
		return
	}

	//update the users
	err = ioutil.WriteFile(".data/users/users.json", upUser, 0666)
	if err != nil {
		return
	}

	w.Write([]byte("User was update"))
	return nil
}
