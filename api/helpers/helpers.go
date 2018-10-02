package helpers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetBody ...
func GetBody(r *http.Request) []byte {
	// read the body content
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	return body
}

//GetQueryID string
func GetQueryID(r *http.Request) (string, error) {
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
func Token(size int) string {
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
