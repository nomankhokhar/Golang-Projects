package routes

import (
	"todo-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initialises the Gin engine and registers all routes
func SetupRouter(handler controllers.TodoHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/todos", handler.GetAllTodos)
		api.GET("/todos/:id", handler.GetTodoByID)
		api.POST("/todos", handler.CreateTodo)
		api.PUT("/todos/:id", handler.UpdateTodo)
		api.DELETE("/todos/:id", handler.DeleteTodo)
	}

	return r
}
