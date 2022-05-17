package repository

import (
	"database/sql"

	"github.com/AsaHero/todolist/db/models"
)

type Authorization interface {
	CreateAccount(user models.Users) (int, error)
	GetUser(username, password string) (models.Users, error)
}

type ToDoList interface {

}

type ToDoItem interface {

}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
	}
}