package models

type Book struct {
	Title       string  `json:"title"`
	Author      *Author `json:"author"`
	Publication string  `json:"publication"`
}

type Author struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
