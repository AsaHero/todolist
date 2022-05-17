package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userTable = "Users"
	todolistTable = "ToDoList"
	todoItemTable = "ToDoItem"
	userslistTable = "UsersList"
	itemlistTable = "ItemList"
)


type Config struct {
	Port string
	Host string
	Useranme string
	Password string
	DBName string
	SSLMode string
}

func NewMysqlDB(conf Config) (*sql.DB, error){
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Useranme, conf.Password, conf.Host, conf.Port, conf.DBName))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}