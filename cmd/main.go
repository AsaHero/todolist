package main

import (
	"log"

	"github.com/AsaHero/todolist"
)

func main() {
	server := new(todolist.Server)
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("error on runnig the server - %s", err.Error())
	}
}
