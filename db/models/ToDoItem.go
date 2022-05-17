package models

type ToDoItem struct {
	Id uint16 `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}