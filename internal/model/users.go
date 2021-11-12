package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	IsActive bool `json:"is_active"`
	Group string `json:"group"`
}

type Group struct {
	ID int `json:"id"`
	Name string `json:"name"`
	IsActive bool `json:"is_active"`
}

