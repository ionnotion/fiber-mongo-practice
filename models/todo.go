package models

type Todo struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Status bool `json:"status"`
}