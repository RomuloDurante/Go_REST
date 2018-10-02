package users

import (
	"io/ioutil"
	"net/http"
)

// Get service
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	// open the user file
	users, err := ioutil.ReadFile(".data/users/users.json")
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "aplication/json")
	w.Write(users)

	return nil
}
