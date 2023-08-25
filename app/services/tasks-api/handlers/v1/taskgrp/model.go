package taskgrp

import (
	"fmt"
	"time"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/sys/validate"
)

// AppTask represents information about an individual task.
type AppTask struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	DueDate     string `json:"dueDate"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppTask(tsk task.Task) AppTask {

	return AppTask{
		ID:          tsk.ID.String(),
		Name:        tsk.Name,
		Description: tsk.Description,
		Type:        tsk.Type,
		Label:       tsk.Label,
		DueDate:     tsk.DueDate.Format(time.RFC3339),
		DateCreated: tsk.DateCreated.Format(time.RFC3339),
		DateUpdated: tsk.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewTask contains information needed to create a new task.
type AppNewTask struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Label       string `json:"label" validate:"required"`
	DueDate     string `json:"dueDate" validate:"required"`
}

func toCoreNewTask(app AppNewTask) (task.NewTask, error) {
	dueDate, err := time.Parse(app.DueDate, time.RFC3339)
	if err != nil {
		fmt.Errorf("parsing date: %w", err)
	}
	tsk := task.NewTask{
		Name:        app.Name,
		Description: app.Description,
		Type:        app.Type,
		Label:       app.Label,
		DueDate:     dueDate,
	}

	return tsk, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewTask) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// AppUpdateTask contains information needed to update a task.
type AppUpdateTask struct {
	Label *string `json:"label" validate:"required"`
}

func toCoreUpdateTask(app AppUpdateTask) (task.UpdateTask, error) {
	// var roles []task.Role
	// if app.Roles != nil {
	// 	roles = make([]task.Role, len(app.Roles))
	// 	for i, roleStr := range app.Roles {
	// 		role, err := task.ParseRole(roleStr)
	// 		if err != nil {
	// 			return task.UpdateTask{}, fmt.Errorf("parsing role: %w", err)
	// 		}
	// 		roles[i] = role
	// 	}
	// }

	nt := task.UpdateTask{
		Label: app.Label,
	}

	return nt, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateTask) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
