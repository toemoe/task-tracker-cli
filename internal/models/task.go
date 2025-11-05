package models

const (
	StatusComplete   = "complete"
	StatusInProgress = "in-progress"
	StatusTodo       = "todo"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
