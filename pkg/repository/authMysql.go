package repository

import (
	"database/sql"
	"fmt"

	"github.com/AsaHero/todolist/db/models"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateAccount(user models.Users) (int, error) {
	var id int64
	
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, HEX(?))", userTable)
	statement, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	res, err := statement.Exec(user.Name, user.Username, user.Password_hash)
	
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(username, password string) (models.Users, error) {
	var user models.Users
	
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=? AND password_hash=HEX(?)", userTable)
	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}
	defer statement.Close()
	row := statement.QueryRow(username, password)
	err = row.Scan(&user.Id, &user.Name, &user.Username, &user.Password_hash)

	return user, err
}