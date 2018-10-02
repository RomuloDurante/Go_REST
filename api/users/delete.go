package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RomuloDurante/WordHunter/api/helpers"
)

// Delete service
func Delete(w http.ResponseWriter, r *http.Request) (err error) {
	// open the user file
	users, err := ioutil.ReadFile(".data/users/users.json")
	if err != nil {
		return
	}
	var checkUser UserData
	json.Unmarshal(users, &checkUser)

	// get the id from url
	intID, err := helpers.GetQueryID(r)
	if err != nil {
		return
	}

	for key, value := range checkUser.Users {
		if intID == value.ID && value.ActiveUser == true {
			checkUser.Users[key].ActiveUser = false
			checkUser.UserInfo.DetetedUsers++
			checkUser.UserInfo.ActiveUsers--
			// create update user
			upUser, err := json.MarshalIndent(checkUser, "", "")
			if err != nil {
				return err
			}

			// update the users
			err = ioutil.WriteFile(".data/users/users.json", upUser, 0666)
			if err != nil {
				return err
			}
			w.Write([]byte("User was deleted !"))

		}

	}

	return nil
}
