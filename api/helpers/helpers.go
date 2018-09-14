package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetBody ...
func GetBody(r *http.Request) []byte {
	// read the body content
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	return body
}

/**********************************/

// Config file ...
type Config struct {
	User user `json:"user"`
	Book book `json:"book"`
}

type user struct {
	ID int
}

type book struct {
	ID int
}

// CheckConfig return the config file
func CheckConfig() Config {
	var conf Config

	file, err := os.Open(".data/config/config.json")

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&conf)

	return conf
}

// UpdateConfig update the config file
func UpdateConfig(c Config) {
	data, err := json.MarshalIndent(c, "", "")
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(".data/config/config.json", data, 0666)
}
