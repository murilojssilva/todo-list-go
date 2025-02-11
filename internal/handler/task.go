package handler

import (
	"net/http"
	"strconv"
	"time"
	"todo-list/internal/data"
	"todo-list/internal/models"
	"todo-list/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"tasks": data.Tasks,
	})
}

func PostTasks(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	if err := service.ValidateTaskTitle(&newTask, data.Tasks); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newTask.ID = len(data.Tasks) + 1
	newTask.Completed = false
	newTask.Created_at = time.Now()
	data.Tasks = append(data.Tasks, newTask)

	data.SaveTask()

	c.JSON(http.StatusCreated, newTask)
}

func GetTasksById(c *gin.Context) {
	idParam := c.Param("id")
	
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": err.Error(),
		})
	return
	}

	for _, p := range data.Tasks {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H {
		"message": "Task not found",
	})	
}

func DeleteTaskById(c *gin.Context){
	idParam := c.Param("id")

    id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": err.Error(),
		})
		return
	}
	for i, p := range data.Tasks {
		if p.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			data.SaveTask()
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})	
}