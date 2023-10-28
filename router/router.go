package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelomoresco/go-todo-crud/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(todoController *controller.TodoController) *gin.Engine{
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", todoController.FindAll)
	tagsRouter.GET("/:todoId", todoController.FindById)
	tagsRouter.POST("", todoController.Create)
	tagsRouter.PUT("/:todoId", todoController.Update)
	tagsRouter.DELETE("/:todoId", todoController.Delete)

	return router
}