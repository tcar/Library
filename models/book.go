package models

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`

	Authors []string `json:"authors"`

	Description string `json:"description"`

	Image string `json:"image"`

	Book_owner int `json:"book_owner"`
}
