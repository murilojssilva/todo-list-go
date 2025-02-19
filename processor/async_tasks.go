package processor

import (
	"fmt"
	"todo-list/internal/models"
)

func PostAsyncTasks(taskChannel <-chan models.Task, done chan<- bool) {
	for task := range taskChannel {
		fmt.Printf("[%s] Tarefa %s | Status: %s\n", 
			task.Created_at.Format("02/Jan/2006 15:04:05"), 
			task.Title, 
			map[bool]string{true: "Concluído"}[task.Completed],
		)
	}
	done <- true
}
