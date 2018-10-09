package controller

import "fmt"

// DataUser ...
type DataUser struct {
	UserInfo infoUsers `json:"userInfo"`
	Users    []User    `json:"users"`
}

// InfoUsers
type infoUsers struct {
	LastID        int `json:"lastId"` // TODO: I don't need the last ID because i will use the toke functions from helpers
	NumberOfUsers int `json:"numberOfUsers"`
	ActiveUsers   int `json:"activeUsers"`
	DetetedUsers  int `json:"detetedUsers "`
}

/***********************************/

// User struct
type User struct {
	ActiveUser bool              `json:"activeUser"`
	ID         string            `json:"id"`
	FirstName  string            `json:"firstName"`
	LastName   string            `json:"lastName"`
	Email      string            `json:"email"`
	Password   string            `json:"password"`
	Agree      bool              `json:"agree"`
	WordVault  map[string]string `json:"wordVault"`
	MyBooks    []myBooks         `json:"myBooks"`
}

//myBooks
type myBooks struct {
	ID       int `json:"id"`
	LastPage int `json:"lastPage"`
	//TODO: notes ?
}

/*************Methods USERDATA*****************************/

//ShowItem ...
func (dataUser DataUser) ShowItem() {
	fmt.Printf("%+v\n", dataUser)
}

// TakeUser ...
func (dataUser DataUser) TakeUser() ([]User, error) {
	return dataUser.Users, nil
}

//ByID ...
func (dataUser DataUser) ByID(item, id string) (interface{}, error) {
	if item == "user" {
		for _, user := range dataUser.Users {
			if user.ID == id {
				return &user, nil
			}
		}
	}

	return nil, fmt.Errorf("%v", "Could not find the data by ID")
}
