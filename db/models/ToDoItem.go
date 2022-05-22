package models

import "errors"

type ToDoItem struct {
	Id uint16 `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

type UpdateToDoItem struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
	Done *bool `json:"done"`
}

func (i *UpdateToDoItem) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("no update values provided")
	}
	return nil
}