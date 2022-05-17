package models

type Users struct {
	Id uint16 `json:"-"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password_hash string `json:"password_hash"`
}