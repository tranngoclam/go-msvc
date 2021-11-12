package model

import (
	"github.com/google/uuid"
	"time"
)

type TaskStatus string

const (
	TaskStatusToDo       = "todo"
	TaskStatusInProgress = "in_progress"
	TaskStatusDone       = "done"
	TaskStatusCanceled   = "canceled"
)

type TaskPriority string

const (
	TaskPriorityLow    = "low"
	TaskPriorityMedium = "medium"
	TaskPriorityHigh   = "high"
)

type Task struct {
	ID          uuid.UUID    `json:"id"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
