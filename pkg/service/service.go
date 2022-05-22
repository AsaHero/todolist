package service

import (
	"github.com/AsaHero/todolist/db/models"
	"github.com/AsaHero/todolist/pkg/repository"
)

type Authorization interface {
	CreateAccount(user models.Users) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ToDoList: NewToDoListService(repo.ToDoList),
		ToDoItem: NewToDoItemService(repo.ToDoItem, repo.ToDoList),
	}
}