package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AsaHero/todolist/db/models"
)

type ToDoItemMysql struct {
	db *sql.DB
}

func NewToDoItemMysql(db *sql.DB) *ToDoItemMysql {
	return &ToDoItemMysql{
		db: db,
	}
}

func (r *ToDoItemMysql) Create(userId int, listId int, item models.ToDoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	queryToDoItem := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES(?, ?, ?)", todoItemTable)
	res, err := tx.Exec(queryToDoItem, item.Title, item.Description, item.Done)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	queryItemList := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES(?, ?)", itemlistTable)
	_, err = tx.Exec(queryItemList, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemId), tx.Commit()

}

func (r *ToDoItemMysql) GetAll(userId, listId int) ([]models.ToDoItem, error) {
	var items []models.ToDoItem

	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s il ON ti.id = il.item_id INNER JOIN %s ul ON il.list_id = ul.list_id WHERE ul.list_id=? AND ul.users_id=?",
		todoItemTable, itemlistTable, userslistTable)
	rows, err := r.db.Query(query, listId, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.ToDoItem
		rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
		items = append(items, item)
	}

	return items, nil
}

func (r *ToDoItemMysql) GetById(userId, itemId int) (models.ToDoItem, error) {
	var item models.ToDoItem

	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s il ON ti.id = il.item_id INNER JOIN %s ul ON il.list_id = ul.list_id WHERE ti.id=? AND ul.users_id=?",
		todoItemTable, itemlistTable, userslistTable)

	row := r.db.QueryRow(query, itemId, userId)
	err := row.Scan(&item.Id, &item.Title, &item.Description, &item.Done)

	return item, err
}

func (r *ToDoItemMysql) DeleteById(userId, itemId int) error {
	query := fmt.Sprintf("DELETE ti, il FROM %s ti INNER JOIN %s il ON ti.id = il.item_id INNER JOIN %s ul ON il.list_id = ul.list_id WHERE ti.id=? AND ul.users_id=?",
		todoItemTable, itemlistTable, userslistTable)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *ToDoItemMysql) Update(userId, itemId int, updateItem models.UpdateToDoItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if updateItem.Title != nil {
		setValues = append(setValues, "title=?")
		args = append(args, updateItem.Title)
	}
	if updateItem.Description != nil {
		setValues = append(setValues, "description=?")
		args = append(args, updateItem.Description)
	}
	if updateItem.Done != nil {
		setValues = append(setValues, "done=?")
		args = append(args, updateItem.Done)
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti INNER JOIN %s il on ti.id = il.item_id INNER JOIN %s ul ON ul.list_id = il.list_id SET %s WHERE ti.id=%d AND ul.users_id=%d",
		todoItemTable, itemlistTable, userslistTable, setQuery, itemId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
