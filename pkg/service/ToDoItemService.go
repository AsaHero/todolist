package service

import (
	"github.com/AsaHero/todolist/db/models"
	"github.com/AsaHero/todolist/pkg/repository"
)

type ToDoItemService struct {
	repoItem repository.ToDoItem
	repoList repository.ToDoList
}

func NewToDoItemService(repoItem repository.ToDoItem, repoList repository.ToDoList) *ToDoItemService {
	return &ToDoItemService{
		repoItem: repoItem,
		repoList: repoList,
	}
}

func (s *ToDoItemService) Create(userId int, listId int, item models.ToDoItem) (int, error) {
	_, err := s.repoList.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repoItem.Create(userId, listId, item)
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]models.ToDoItem, error) {
	_, err := s.repoList.GetById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repoItem.GetAll(userId, listId)
}

func (s *ToDoItemService) GetById(userId, itemId int) (models.ToDoItem, error) {
	return s.repoItem.GetById(userId, itemId)
}

func (s *ToDoItemService) DeleteById(userId, itemId int) error {
	return s.repoItem.DeleteById(userId, itemId)
}

func (s *ToDoItemService) Update(userId, itemId int, updateItem models.UpdateToDoItem) error {
	err := updateItem.Validate()
	if err != nil {
		return err
	}
	return s.repoItem.Update(userId, itemId, updateItem)
}