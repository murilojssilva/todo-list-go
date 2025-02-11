package main

import (
	"todo-list/internal/handler"

	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasks/:id", handler.GetTasksById)
	router.POST("/tasks", handler.PostTasks)
	router.DELETE("/pizzas/:id", handler.DeleteTaskById)

	router.Run()
}