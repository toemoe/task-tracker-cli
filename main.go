package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID     int
	Name   string
	Status string
}

func addTask() {
	fmt.Println("Add task")
}

func deleteTask() {
	fmt.Println("Delete task")
}

func updateTask() {
	fmt.Println("Update task")
}

func changeMarkTask(status string) {
	switch status {
	case "in-progress":
		fmt.Println("Mark task as in-progress")
	case "mark-done":
		fmt.Println("Mark task as done")
	default:
		fmt.Println("Invalid status")
	}
}

func getListTasks(status string) {
	switch status {
	case "done":
		fmt.Println("List done tasks")
	case "todo":
		fmt.Println("List todo tasks")
	case "in-progress":
		fmt.Println("List in-progress tasks")
	default:
		fmt.Println("List all tasks")
	}
}

func taskAction(action string, args ...string) {
	switch action {
	case "add":
		addTask()
	case "delete":
		deleteTask()
	case "update":
		updateTask()
	case "list":
		getListTasks("")
	case "mark":
		changeMarkTask("")
	default:
		fmt.Println("Invalid action")
	}
}

func main() {
	fmt.Println("Welcome to Task Manager")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter action:")
		if !scanner.Scan() {
			fmt.Println("Error reading input")
			break
		}
		action := scanner.Text()
		action = strings.ToLower(action)

		taskAction(action)
	}

}
