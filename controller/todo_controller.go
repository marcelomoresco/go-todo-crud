package controller

import (
	"net/http"
	"strconv"

	"github.com/marcelomoresco/go-todo-crud/data/request"
	"github.com/marcelomoresco/go-todo-crud/data/response"
	helper "github.com/marcelomoresco/go-todo-crud/helpers"
	service "github.com/marcelomoresco/go-todo-crud/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController{
	return &TodoController{
		todoService: service,
	}
}

func(controller *TodoController) Create(ctx *gin.Context) {
	createTodoRequest := request.CreateTodoRequest{}
	err := ctx.ShouldBindJSON(&createTodoRequest)
	helper.ErrorPanic(err)
	controller.todoService.Create(createTodoRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "OK",
		Data: nil,
	}
	ctx.Header("Content-Type","application-json")
	ctx.JSON(http.StatusOK,webResponse)
}

func (controller *TodoController) Delete(ctx *gin.Context){
	todoId := ctx.Param("todoId")
	id,err := strconv.Atoi(todoId);
	helper.ErrorPanic(err)
	controller.todoService.Delete(id)
	webResponse := response.Response{
		Code: http.StatusNoContent,
		Status: "No content",
		Data: nil,
	}
	ctx.Header("Content-Type","application-json")
	ctx.JSON(http.StatusNoContent,webResponse)
}

func (controller *TodoController) FindById(ctx *gin.Context){
	todoId := ctx.Param("todoId")
	id,err := strconv.Atoi(todoId);
	helper.ErrorPanic(err)
	todoResponse := controller.todoService.FindById(id)
	webResponse := response.Response{
		Code: http.StatusNoContent,
		Status: "No content",
		Data: todoResponse,
	}
	ctx.Header("Content-Type","application-json")
	ctx.JSON(http.StatusNoContent,webResponse)
}
func (controller *TodoController) Update(ctx *gin.Context){
	updateTodoRequest := request.UpdateTodoRequest{}
	err := ctx.ShouldBindJSON(&updateTodoRequest)
	helper.ErrorPanic(err)

	todoId := ctx.Param("todoId")
	id,err := strconv.Atoi(todoId)
	helper.ErrorPanic(err)
	updateTodoRequest.Id = id

	controller.todoService.Update(updateTodoRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "OK",
		Data: nil,
	}
	ctx.Header("Content-Type","application-json")
	ctx.JSON(http.StatusOK,webResponse)
}
func (controller *TodoController) FindAll(ctx *gin.Context){
	todoResponse := controller.todoService.FindAll()
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "OK",
		Data: todoResponse,
	}
	ctx.Header("Content-Type","application-json")
	ctx.JSON(http.StatusOK,webResponse)
}