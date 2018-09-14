package books

import (
	"net/http"
)

// Put service
func Put(w http.ResponseWriter, r *http.Request) (err error) {
	// var user User // we need 2 users structs to compare
	// var updatedUser User

	// // get the query string
	// query := r.URL.Query()
	// var id string

	// for key, value := range query {
	// 	if key == "id" {
	// 		for _, queryID := range value {
	// 			id = queryID
	// 		}
	// 	}
	// }

	// // try use id to read user
	// data, err := ioutil.ReadFile(".data/" + id + ".json")

	// if err != nil {
	// 	return
	// }

	// // if ok push the data into a struct
	// json.Unmarshal(data, &user)

	// // read the body content
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)

	// // push the body into newUser
	// err = json.Unmarshal(body, &updatedUser)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// // check if the data send to server matches with data is already there, if not update the data
	// for {
	// 	if user.FirstName != updatedUser.FirstName {
	// 		user.FirstName = updatedUser.FirstName
	// 		continue
	// 	} else if user.LastName != updatedUser.LastName {
	// 		user.LastName = updatedUser.LastName
	// 		continue
	// 	} else if user.Email != updatedUser.Email {
	// 		user.Email = updatedUser.Email
	// 	}
	// 	break
	// }

	// // create update user
	// upUser, err := json.Marshal(user)
	// if err != nil {
	// 	return
	// }

	// ioutil.WriteFile(id+".json", upUser, 0666)
	// os.Rename(id+".json", ".data/"+id+".json")

	// w.Write([]byte("User was update"))
	w.Write([]byte("books put works"))
	return nil
}
