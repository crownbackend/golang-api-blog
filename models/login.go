package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
