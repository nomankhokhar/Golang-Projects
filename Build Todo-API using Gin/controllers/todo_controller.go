package controllers

import (
	"net/http"
	"strconv"

	"todo-crud/models"

	"github.com/gin-gonic/gin"
)

// TodoHandler defines the interface for todo operations
type TodoHandler interface {
	GetAllTodos(c *gin.Context)
	GetTodoByID(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

// TodoController holds a reference to the store
type TodoController struct {
	Store  *models.TodoStore
	Number int
}

// NewTodoController creates a controller with the given store
func NewTodoController(store *models.TodoStore) *TodoController {
	return &TodoController{Store: store, Number: 0}
}

// GetAllTodos returns every todo
func (tc *TodoController) GetAllTodos(c *gin.Context) {
	tc.Number++
	c.JSON(http.StatusOK, tc.Store.GetAll())
}

// GetTodoByID returns a single todo by its ID
func (tc *TodoController) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo := tc.Store.GetByID(id)
	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	tc.Number++
	c.JSON(http.StatusOK, todo)
}

// CreateTodo creates a new todo
func (tc *TodoController) CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := tc.Store.Create(newTodo)

	tc.Number++
	c.JSON(http.StatusCreated, created)
}

// UpdateTodo updates an existing todo by ID
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := tc.Store.Update(id, updatedTodo)
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	tc.Number++
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}

// DeleteTodo removes a todo by ID
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if !tc.Store.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	tc.Number++
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
