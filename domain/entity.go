package domain

// Define the Structs

type Article struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Date   string  `json:"publishedAt"`
	Medias []Media `json:"medias"`
}

type Media struct {
	ID          int    `json:"id"`
	ContentURL  string `json:"contentUrl"`
	ContentType string `json:"contentType"`
}
