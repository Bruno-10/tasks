// Package task provides an example of a core business API. Right now these
// calls are just wrapping the data/data layer. But at some point you will
// want auditing or something that isn't specific to the data/store layer.
package task

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	Query(ctx context.Context) ([]Task, error)
	Count(ctx context.Context) (int, error)
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

// Query retrieves a list of existing tasks from the database.
func (c *Core) Query(ctx context.Context) ([]Task, error) {
	tasks, err := c.storer.Query(ctx)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return tasks, nil
}

// Count returns the total number of tasks in the store.
func (c *Core) Count(ctx context.Context) (int, error) {
	return c.storer.Count(ctx)
}
