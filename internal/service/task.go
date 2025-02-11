package service

import (
	"errors"
	"todo-list/internal/models"
)

func ValidateTaskTitle(task *models.Task, existingTasks []models.Task) error {
	for _, t := range existingTasks {
		if t.Title == task.Title {
			return errors.New("já há uma tarefa com esse título")
		}
	}
	return nil
}
