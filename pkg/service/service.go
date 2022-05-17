package service

import (
	"github.com/AsaHero/todolist/db/models"
	"github.com/AsaHero/todolist/pkg/repository"
)

type Authorization interface {
	CreateAccount(user models.Users) (int, error)
	GenerateToken(username, password string) (string, error)
}

type ToDoList interface {

}

type ToDoItem interface {

}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}