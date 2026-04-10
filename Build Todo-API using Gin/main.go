package main

import (
	"todo-crud/controllers"
	"todo-crud/models"
	"todo-crud/routes"
)

func main() {
	store := models.NewTodoStore()
	tc := controllers.NewTodoController(store)
	r := routes.SetupRouter(tc)
	r.Run(":8080")
}
