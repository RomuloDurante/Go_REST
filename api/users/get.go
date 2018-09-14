package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Get service
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	// open the user file
	users, err := ioutil.ReadFile(".data/users/users.json")
	if err != nil {
		return
	}
	var checkUser UserData
	json.Unmarshal(users, &checkUser)

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
	var showUser User
	for _, user := range checkUser.Users {
		if intID == user.ID && user.ActiveUser == true {
			user.Password = ""
			showUser = user

			// get the user data
			data, err := json.MarshalIndent(showUser, "", "")
			if err != nil {
				return err
			}

			w.Header().Set("Content-type", "aplication/json")
			w.Write(data)

		}

	}

	w.Write([]byte("This user was deleted"))
	return nil
}
