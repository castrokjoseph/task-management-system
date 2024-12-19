package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string
type TaskPriority string

const (
	StatusTodo        TaskStatus = "TODO"
	StatusInProgress  TaskStatus = "IN_PROGRESS"
	StatusDone        TaskStatus = "DONE"
	StatusBlocked     TaskStatus = "BLOCKED"

	PriorityLow       TaskPriority = "LOW"
	PriorityMedium    TaskPriority = "MEDIUM"
	PriorityHigh      TaskPriority = "HIGH"
	PriorityCritical  TaskPriority = "CRITICAL"
)

type Task struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	Title       string        `json:"title" db:"title"`
	Description string        `json:"description" db:"description"`
	Status      TaskStatus    `json:"status" db:"status"`
	Priority    TaskPriority  `json:"priority" db:"priority"`
	UserID      uuid.UUID     `json:"user_id" db:"user_id"`
	ProjectID   *uuid.UUID    `json:"project_id,omitempty" db:"project_id"`
	Deadline    *time.Time    `json:"deadline,omitempty" db:"deadline"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

type TaskCreate struct {
	Title       string        `json:"title" validate:"required"`
	Description string        `json:"description"`
	Status      TaskStatus    `json:"status"`
	Priority    TaskPriority  `json:"priority"`
	ProjectID   *uuid.UUID    `json:"project_id"`
	Deadline    *time.Time    `json:"deadline"`
}