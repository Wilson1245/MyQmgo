package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
