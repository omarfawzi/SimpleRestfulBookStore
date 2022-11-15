package models

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title" validate:"required" faker:"name"`
	Author string `json:"author" validate:"required" faker:"name"`
}
