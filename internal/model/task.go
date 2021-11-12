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
	ID          uuid.UUID
	Description string
	Priority    TaskPriority
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
