package service

import (
	"github.com/marcelomoresco/go-todo-crud/data/request"
	"github.com/marcelomoresco/go-todo-crud/data/response"
	helper "github.com/marcelomoresco/go-todo-crud/helpers"
	"github.com/marcelomoresco/go-todo-crud/model"
	repositories "github.com/marcelomoresco/go-todo-crud/repository"

	"github.com/go-playground/validator"
)

type TodoServiceImpl struct {
	TodoRepository repositories.TodoRepository
	Validate       *validator.Validate
}

// Create implements TodoService.
func (t *TodoServiceImpl) Create(todo request.CreateTodoRequest) {
	err := t.Validate.Struct(todo)
	helper.ErrorPanic(err)
	todoModel := model.Todo{
		Name: todo.Name,
		Description: todo.Description,
		Done: todo.Done,
	}
	t.TodoRepository.Save(todoModel)
}

// Delete implements TodoService.
func (t *TodoServiceImpl) Delete(todoId int) {
	t.TodoRepository.Delete(todoId)
}

// FindAll implements TodoService.
func (t *TodoServiceImpl) FindAll() []response.TodoResponse {
	result := t.TodoRepository.FindAll()

	var tags []response.TodoResponse
	for _, value := range result {
		tag := response.TodoResponse{
			Id:   value.Id,
			Name: value.Name,
			Description: value.Description,
			Done: value.Done,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TodoService.
func (t *TodoServiceImpl) FindById(todoId int) response.TodoResponse {
	todoData, err := t.TodoRepository.FindById(todoId)
	helper.ErrorPanic(err)

	todoResponse := response.TodoResponse{
		Id:   todoData.Id,
		Name: todoData.Name,
		Description: todoData.Description,
		Done: todoData.Done,
	}
	return todoResponse
}

// Update implements TodoService.
func (t *TodoServiceImpl) Update(todo request.UpdateTodoRequest) {
	todoData, err := t.TodoRepository.FindById(todo.Id)
	helper.ErrorPanic(err)
	todoData.Name = todo.Name
	todoData.Description = todo.Description
	todoData.Done = todo.Done
	t.TodoRepository.Update(todoData)
}

func NewTodoServiceImpl(todoRepository repositories.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		Validate:       validate,
	}
}
