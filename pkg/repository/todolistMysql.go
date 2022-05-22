package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AsaHero/todolist/db/models"
)

type ToDoListMysql struct {
	db *sql.DB
}

func NewToDoListMysql(db *sql.DB) *ToDoListMysql {
	return &ToDoListMysql{
		db: db,
	}
}

func (r *ToDoListMysql) Create(userId int, list models.ToDoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	createListquery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES(?, ?)", todolistTable)
	res, err := tx.Exec(createListquery, list.Title, list.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	
	listId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersList := fmt.Sprintf("INSERT INTO %s (users_id, list_id) VALUES(?, ?)", userslistTable)
	_, err = tx.Exec(createUsersList, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	
	return int(listId), tx.Commit()

}

func (r *ToDoListMysql) GetAll(userId int) ([]models.ToDoList, error) {
	var lists []models.ToDoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul WHERE ul.id = ?", todolistTable, userslistTable)
	rows, err := r.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var list models.ToDoList
		rows.Scan(&list.Id, &list.Title, &list.Description)
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *ToDoListMysql) GetById(userId int, listId int) (models.ToDoList, error) {
	var list models.ToDoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul WHERE ul.users_id = ? AND tl.id = ?", todolistTable, userslistTable)
	row  := r.db.QueryRow(query, userId, listId)
	
	err := row.Scan(&list.Id, &list.Title, &list.Description)

	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *ToDoListMysql) DeleteById(userId int, listId int) error {
	deleteQuery := fmt.Sprintf("DELETE tl, ul FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.users_id=? AND ul.list_id=?", todolistTable, userslistTable)
	_, err := r.db.Exec(deleteQuery, userId, listId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ToDoListMysql) Update(userId int, listId int, updateList models.UpdateToDoList) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if updateList.Title != nil {
		setValues = append(setValues, "title=?")
		args = append(args, updateList.Title)
	}
	if updateList.Description != nil {
		setValues = append(setValues, "description=?")
		args = append(args, updateList.Description)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s AS tl INNER JOIN %s AS ul ON tl.id = ul.list_id SET %s WHERE ul.users_id=%d AND ul.list_id=%d", todolistTable, userslistTable, setQuery, userId, listId)

	_, err := r.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}
