package books

// BookData ...
type BookData struct {
	BookInfo infoBook `json:"bookInfo"`
	Books    []Book   `json:"books"`
}

/************************/
// infobook
type infoBook struct {
	LastID        int         `json:"lastId"`
	NumberOfBooks int         `json:"numberOfBooks"`
	Rating        []bestBooks `json:"rating"`
}

// bestBooks
type bestBooks struct {
	ID     int    `json:"id"`
	Rating int    `json:"rating"`
	Name   string `json:"name"`
}

/************************/

// Book ...
type Book struct {
	ID        int    `json:"id"`
	Rating    int    `json:"rating"`
	Name      string `json:"name"`
	Author    author `json:"author"`
	AboutBook string `json:"aboutBook"`
	Dificulty int    `json:"dificulty"`
}

type author struct {
	Name        string `json:"name"`
	AboutAuthor string `json:"aboutAuthor"`
}
