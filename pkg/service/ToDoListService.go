package service

import (
	"github.com/AsaHero/todolist/db/models"
	"github.com/AsaHero/todolist/pkg/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{
		repo: repo,
	}
}

func (s *ToDoListService) Create(userId int, list models.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ToDoListService) GetAll(userId int) ([]models.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func (s *ToDoListService) GetById(userId int, listId int) (models.ToDoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ToDoListService) DeleteById(userId int, listId int) error {
	return s.repo.DeleteById(userId, listId)
}

func (s *ToDoListService) Update(userId int, listId int, updateList models.UpdateToDoList) error {
	if err := updateList.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, updateList)
}