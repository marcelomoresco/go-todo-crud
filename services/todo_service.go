package service

import (
	"github.com/marcelomoresco/go-todo-crud/data/request"
	"github.com/marcelomoresco/go-todo-crud/data/response"
)

type TodoService interface {
	Create(todo request.CreateTodoRequest)
	Update(todo request.UpdateTodoRequest)
	Delete(todoId int)
	FindById(todoId int) response.TodoResponse
	FindAll() []response.TodoResponse
}