package helpers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Helpers ...
type Helpers struct{}

//Help ...-> helpers interface
type Help interface {
	ByID(item, id string) (interface{}, error)
}

//GetQueryID string
func (Helpers) GetQueryID(r *http.Request) (string, error) {
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

	return id, nil

}

// Token ...
func (Helpers) Token(size int) string {
	var token string

	for i := 0; i < size; i++ {
		// To produce varying sequences, give it a seed that changes
		seed := rand.NewSource(time.Now().UnixNano())
		r := rand.New(seed)

		//use the new rand to produce randon numbers
		random1 := r.Intn(25)
		random2 := r.Intn(25)
		random3 := r.Intn(10)

		// use the ASCII to generate token
		if i%2 == 0 {
			token += string(97+random1) + string(65+random2) + strconv.Itoa(random3)
		} else {
			token += string(97+random1) + string(65+random2) + strconv.Itoa(random3)
		}
	}

	return token

}

//Capitalize words
func (Helpers) Capitalize(s string) string {
	cap := []rune(s)
	upper := strings.ToUpper(string(cap[0]))
	cap[0] = []rune(upper)[0]

	upper = string(cap)

	return upper
}

//JSONstring ...
func (Helpers) JSONstring(dt interface{}, err error) ([]byte, error) {

	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(dt, " ", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

//JSONstruct ...
func (Helpers) JSONstruct(dt []byte, str Help) (Help, error) {

	err := json.Unmarshal(dt, &str)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}

	return str, nil
}
