package main

import (
	"todo-list/internal/data"
	"todo-list/internal/handler"

	"github.com/gin-gonic/gin"
)
func main() {
	data.LoadTasks()
	router := gin.Default()
	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasks/:id", handler.GetTasksById)
	router.POST("/tasks", handler.PostTasks)
	router.DELETE("/pizzas/:id", handler.DeleteTaskById)

	router.POST("/tasks/:id/process", handler.ProcessTask)

	router.Run()
}