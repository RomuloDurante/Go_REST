package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RomuloDurante/WordHunter/api/helpers"
)

// Post service
func Post(w http.ResponseWriter, r *http.Request) (err error) {
	// open the data users
	users, err := ioutil.ReadFile(".data/users/users.json")

	if err != nil {
		return
	}

	// create and unmarshal the datauser
	var dataUser UserData
	json.Unmarshal(users, &dataUser)

	// create newUser
	var newUser User

	// create the new id
	id := dataUser.UserInfo.LastID + 1
	// push the new id to config file
	dataUser.UserInfo.LastID = id

	// use the new id to create newUser
	newUser.ID = helpers.Token(10) + strconv.Itoa(id)
	newUser.ActiveUser = true

	// read the body content
	body := helpers.GetBody(r)

	// push the body into newUser
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// push newUser to dataUser
	dataUser.Users = append(dataUser.Users, newUser)
	dataUser.UserInfo.NumberOfUsers++
	dataUser.UserInfo.ActiveUsers++
	// parse the newUser to json
	data, err := json.MarshalIndent(dataUser, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// use the id to create new user
	err = ioutil.WriteFile(".data/users/users.json", data, 0666)
	if err != nil {
		return
	}

	w.Write([]byte("User was create"))
	return nil
}
