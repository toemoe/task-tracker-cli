package models

const (
	StatusComplete   = "complete"
	StatusInProgress = "in-progress"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
