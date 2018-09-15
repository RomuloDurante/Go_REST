package users

// UserData ...
type UserData struct {
	UserInfo infoUsers `json:"userInfo"`
	Users    []User    `json:"users"`
}

// InfoUsers
type infoUsers struct {
	LastID        int `json:"lastId"`
	NumberOfUsers int `json:"numberOfUsers"`
	ActiveUsers   int `json:"activeUsers"`
	DetetedUsers  int `json:"detetedUsers "`
}

// User struct
type User struct {
	ActiveUser bool              `json:"activeUser"`
	ID         int               `json:"id"`
	FirstName  string            `json:"firstName"`
	LastName   string            `json:"lastName"`
	Email      string            `json:"email"`
	Password   string            `json:"password"`
	Agree      bool              `json:"agree"`
	WordVault  map[string]string `json:"wordVault"`
}
