package main

import (
	"fmt"
	"os"

	"github.com/toemoe/task-tracker-cli/internal/handlers"
	"github.com/toemoe/task-tracker-cli/internal/storage"
	"github.com/toemoe/task-tracker-cli/pkg/utils"
)

func main() {
	const filename = "tasks.json"

	if err := storage.LoadTasksFromFile(filename); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli add|delete|update|mark|list [task_id] [task_name]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add [task_name]")
			return
		}
		name := os.Args[2]
		handlers.AddTask(name)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete [task_id]")
			return
		}
		id := os.Args[2]
		if !utils.IsValidUUID(id) {
			fmt.Println("Invalid task ID")
			return
		}
		handlers.DeleteTask(id)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update [task_id] [task_name]")
			return
		}
		id := os.Args[2]
		if !utils.IsValidUUID(id) {
			fmt.Println("Invalid task ID")
			return
		}
		name := os.Args[3]
		handlers.UpdateTask(id, name)
	case "mark":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark [task_id]")
			return
		}
		id := os.Args[2]
		if !utils.IsValidUUID(id) {
			fmt.Println("Invalid task ID")
			return
		}
		handlers.MarkTask(id)
	case "list":
		handlers.ListTasks("")
	case "list-complete":
		handlers.ListTasks("complete")
	case "list-in-progress":
		handlers.ListTasks("in-progress")
	case "list-todo":
		handlers.ListTasks("todo")
	default:
		fmt.Println("Invalid command")
	}

	if err := storage.SaveTasksToFile(filename); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}
