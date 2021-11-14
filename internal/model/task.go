package model

import (
	"time"
)

// TaskStatus defines type alias for the status of a task
type TaskStatus string

// TaskStatus enumeration
const (
	TaskStatusToDo       = "todo"
	TaskStatusInProgress = "in_progress"
	TaskStatusDone       = "done"
	TaskStatusCanceled   = "canceled"
)

// TaskPriority defines type alias for the priority of a task
type TaskPriority string

// TaskPriority enumeration
const (
	TaskPriorityLow    = "low"
	TaskPriorityMedium = "medium"
	TaskPriorityHigh   = "high"
)

// Task defines properties for a task
type Task struct {
	ID          uint64       `json:"id"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
