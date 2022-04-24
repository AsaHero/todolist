package repository

import "database/sql"

type Authorization struct {

}

type ToDoList struct {

}

type ToDoItem struct {

}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}