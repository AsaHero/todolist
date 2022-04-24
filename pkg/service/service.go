package service

import "github.com/AsaHero/todolist/pkg/repository"

type Authorization struct {

}

type ToDoList struct {

}

type ToDoItem struct {

}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}