package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/marcelomoresco/go-todo-crud/config"
	"github.com/marcelomoresco/go-todo-crud/controller"
	error "github.com/marcelomoresco/go-todo-crud/helpers"
	"github.com/marcelomoresco/go-todo-crud/model"
	repositories "github.com/marcelomoresco/go-todo-crud/repository"
	"github.com/marcelomoresco/go-todo-crud/router"
	service "github.com/marcelomoresco/go-todo-crud/services"
)

func main(){


	//Database
	db :=config.DatabaseConfig()
	validate := validator.New()
	db.Table("todos").AutoMigrate(&model.Todo{})

	todoRepository :=repositories.NewTodoRepositoryImpl(db)
	todoService := service.NewTodoServiceImpl(todoRepository,validate)

	todoController := controller.NewTodoController(todoService)

	routes :=router.NewRouter(todoController)


	server:=&http.Server{
		Addr: ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe();
	error.ErrorPanic(err)
}