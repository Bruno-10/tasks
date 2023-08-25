package task

import (
	"time"

	"github.com/google/uuid"
)

// Task represents information about an individual task.
type Task struct {
	ID          uuid.UUID
	Name        string
	Description string
	Type        string
	Label       string
	DueDate     time.Time
	DateCreated time.Time
	DateUpdated time.Time
}

// NewTask contains information needed to create a new task.
type NewTask struct {
	Name        string
	Description string
	Type        string
	Label       string
	DueDate     time.Time
}

// UpdateTask contains information needed to update a task.
type UpdateTask struct {
	Label *string
}
