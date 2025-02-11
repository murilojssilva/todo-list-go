package data

import (
	"encoding/json"
	"todo-list/internal/models"

	"fmt"
	"os"
)

var Tasks []models.Task

func LoadTasks() {
	file, err := os.Open("data/tasks.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Tasks); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SaveTask() {
	file, err := os.Create("data/tasks.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(Tasks); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}