package repositories

import (
	"errors"

	"github.com/marcelomoresco/go-todo-crud/data/request"
	helper "github.com/marcelomoresco/go-todo-crud/helpers"
	"github.com/marcelomoresco/go-todo-crud/model"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TodoRepository.
func (t *TodoRepositoryImpl) Delete(todoId int) {
	var todo model.Todo
	result :=t.Db.Where("id=?",todoId).Delete(&todo)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TodoRepository.
func (t *TodoRepositoryImpl) FindAll() []model.Todo {
	var todos []model.Todo
	result :=t.Db.Find(&todos)
	helper.ErrorPanic(result.Error)
	return todos;
}

// FindById implements TodoRepository.
func (t *TodoRepositoryImpl) FindById(todoId int) (todos model.Todo, err error){
	var todo model.Todo
	result := t.Db.Find(&todo,todoId)
	if result != nil{
		return todo,nil
	}else{
		return todo, errors.New("todo not found")
	}
}

// Save implements TodoRepository.
func (t *TodoRepositoryImpl) Save(todo model.Todo) {
	result :=t.Db.Create(&todo)
	helper.ErrorPanic(result.Error)
}

// Update implements TodoRepository.
func (t *TodoRepositoryImpl) Update(todo model.Todo) {
	var updateTodo = request.UpdateTodoRequest{
		Id:todo.Id,
		Name:todo.Name,
		Description:todo.Description,
		Done:todo.Done,
	}
	result :=t.Db.Model(&todo).Updates(updateTodo)
	helper.ErrorPanic(result.Error)
}

func NewTodoRepositoryImpl(Db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}
