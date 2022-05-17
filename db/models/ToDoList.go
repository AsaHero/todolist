package models

type ToDoList struct {
	Id uint16 `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
}