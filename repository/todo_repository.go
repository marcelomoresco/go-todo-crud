package repositories

import (
	"github.com/marcelomoresco/go-todo-crud/model"
)

type TodoRepository interface {
	Save(todo model.Todo)
	Update(todo model.Todo)
	Delete(todoId int)
	FindById(todoId int)(todos model.Todo, err error)
	FindAll() []model.Todo
}