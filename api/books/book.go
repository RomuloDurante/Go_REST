package books

// Book ...
type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    author `json:"author"`
	About     string `json:"about"`
	Dificulty int    `json:"dificulty"`
}

type author struct {
	Name  string `json:"name"`
	About string `json:"about"`
}
