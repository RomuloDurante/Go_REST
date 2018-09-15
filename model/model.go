package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/RomuloDurante/WordHunter/restApi_GO/api/users"
)

// GetData ...
func GetData() *users.User {
	data, err := http.Get("http://localhost:3000/api/book/?id=1")
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	var user users.User
	json.Unmarshal(body, &user)

	return &user
}
