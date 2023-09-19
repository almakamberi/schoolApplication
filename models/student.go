package models

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"Name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}
