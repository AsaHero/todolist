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
	Create(int, models.ToDoList) (int, error)
	GetAll(int) ([]models.ToDoList, error)
	GetById(int, int) (models.ToDoList, error)
	DeleteById(int, int) error
	Update(int, int, models.UpdateToDoList) error
}

type ToDoItem interface {
	Create(int, int, models.ToDoItem) (int, error)
	GetAll(int, int) ([]models.ToDoItem, error)
	GetById(int, int) (models.ToDoItem, error)
	DeleteById(int, int) error
	Update(int, int, models.UpdateToDoItem) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		ToDoList: NewToDoListMysql(db),
		ToDoItem: NewToDoItemMysql(db),
	}
}