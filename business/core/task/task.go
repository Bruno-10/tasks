// Package task provides an example of a core business API. Right now these
// calls are just wrapping the data/data layer. But at some point you will
// want auditing or something that isn't specific to the data/store layer.
package task

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Bruno-10/tasks/business/data/order"
	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("task not found")
	ErrUniqueEmail           = errors.New("label is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, tsk Task) error
	Update(ctx context.Context, tsk Task) error
	Delete(ctx context.Context, tsk Task) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Task, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, taskID uuid.UUID) (Task, error)
	QueryByIDs(ctx context.Context, taskID []uuid.UUID) ([]Task, error)
	QueryByLabel(ctx context.Context, label string) (Task, error)
}

// Core manages the set of APIs for task access.
type Core struct {
	storer Storer
}

// NewCore constructs a core for task api access.
func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

// Create inserts a new task into the database.
func (c *Core) Create(ctx context.Context, nt NewTask) (Task, error) {
	now := time.Now()

	tsk := Task{
		ID:          uuid.New(),
		Name:        nt.Name,
		Description: nt.Description,
		Type:        nt.Type,
		Label:       nt.Label,
		DueDate:     nt.DueDate,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, tsk); err != nil {
		return Task{}, fmt.Errorf("create: %w", err)
	}

	return tsk, nil
}

// Update replaces a task document in the database.
func (c *Core) Update(ctx context.Context, tsk Task, ut UpdateTask) (Task, error) {
	if ut.Label != nil {
		tsk.Label = *ut.Label
	}
	tsk.DateUpdated = time.Now()

	if err := c.storer.Update(ctx, tsk); err != nil {
		return Task{}, fmt.Errorf("update: %w", err)
	}

	return tsk, nil
}

// Delete removes a task from the database.
func (c *Core) Delete(ctx context.Context, tsk Task) error {
	if err := c.storer.Delete(ctx, tsk); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing tasks from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Task, error) {
	tasks, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return tasks, nil
}

// Count returns the total number of tasks in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}

// QueryByID gets the specified task from the database.
func (c *Core) QueryByID(ctx context.Context, taskID uuid.UUID) (Task, error) {
	task, err := c.storer.QueryByID(ctx, taskID)
	if err != nil {
		return Task{}, fmt.Errorf("query: taskID[%s]: %w", taskID, err)
	}

	return task, nil
}

// QueryByLabel gets the specified task from the database by label.
func (c *Core) QueryByLabel(ctx context.Context, label string) (Task, error) {
	task, err := c.storer.QueryByLabel(ctx, label)
	if err != nil {
		return Task{}, fmt.Errorf("query: label[%s]: %w", label, err)
	}

	return task, nil
}
