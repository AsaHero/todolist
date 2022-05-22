package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AsaHero/todolist"
	"github.com/AsaHero/todolist/goenv"
	"github.com/AsaHero/todolist/pkg/handler"
	"github.com/AsaHero/todolist/pkg/repository"
	"github.com/AsaHero/todolist/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initViper(); err != nil {
		log.Fatalf("Error on reading config files: %s\n", err.Error())
	}

	env, err := goenv.Load(".env")
	if err != nil {
		log.Fatalf("Error on reading .env file: %s\n", err.Error())
	}
	
	db, err := repository.NewMysqlDB(repository.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Useranme: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: env.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("error on connecting to the database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	server := new(todolist.Server)
	go func ()  {
		if err := server.Run(viper.GetString("port"), handler.InitRouter()); err != nil {
			log.Fatalf("error on runnig the server - %s", err.Error())
		}
	}()
	log.Println("ToDoApp Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("ToDoApp Server Shutting Down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on ToDoApp server shutting down: %s", err.Error())
	}  
	if err := db.Close(); err != nil {
		log.Fatalf("error occured on closing db connection")
	}

}			



func initViper() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
