package models

import "time"

const (
	StatusComplete   = "complete"
	StatusInProgress = "in-progress"
	StatusTodo       = "todo"
)

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
