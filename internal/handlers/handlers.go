package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/toemoe/task-tracker-cli/internal/models"
	"github.com/toemoe/task-tracker-cli/internal/storage"
)

func AddTask(name string) {
	task := models.Task{ID: uuid.New().String(), Name: name, Status: models.StatusInProgress}
	storage.AddTask(task)

	if _, err := json.Marshal(task); err != nil {
		fmt.Println("serialization error:", err)
		return
	}

	fmt.Printf("task %s added\n", task.Name)
}

func DeleteTask(id string) {
	if storage.DeleteTask(id) {
		fmt.Println("task deleted")
	} else {
		fmt.Println("task not found")
	}
}

func UpdateTask(id string, name string) {
	if storage.UpdateTask(id, name) {
		fmt.Println("task updated successfully")
	} else {
		fmt.Println("task not found")
	}
}

func MarkTask(id string) {
	if storage.MarkTask(id) {
		fmt.Println("status updated")
	} else {
		fmt.Println("task not found")
	}
}

func ListTasks(status string) {
	for _, task := range storage.GetTasks() {
		if status == "" || task.Status == status {
			fmt.Printf("name: %s, status: %s\n", task.Name, task.Status)
		}
	}
}
