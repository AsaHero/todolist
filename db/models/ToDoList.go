package models

import "errors"

type ToDoList struct {
	Id uint16 `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UpdateToDoList struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateToDoList) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("no update date provided")
	}

	return nil
}