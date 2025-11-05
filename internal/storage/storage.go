package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/toemoe/task-tracker-cli/internal/models"
	"github.com/toemoe/task-tracker-cli/pkg/utils"
)

var tasks []models.Task

func LoadTasksFromFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, &tasks)
}

func SaveTasksToFile(filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling tasks: %v", err)
	}
	return os.WriteFile(filename, data, 0644)
}

func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func DeleteTask(id string) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateTask(id, name, description string) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = name
			tasks[i].Description = description
			tasks[i].UpdatedAt = utils.GetTimeNow()
			return true
		}
	}
	return false
}

func MarkTask(id string) bool {
	for i, task := range tasks {
		if task.ID == id {
			switch task.Status {
			case models.StatusTodo:
				tasks[i].Status = models.StatusInProgress
			case models.StatusInProgress:
				tasks[i].Status = models.StatusComplete
			case models.StatusComplete:
				tasks[i].Status = models.StatusTodo
			}
			return true
		}
	}
	return false
}

func GetTasks() []models.Task {
	return tasks
}
