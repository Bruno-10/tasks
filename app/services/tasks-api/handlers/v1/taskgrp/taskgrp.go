// Package taskgrp maintains the group of handlers for task access.
package taskgrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Bruno-10/tasks/business/core/task"
	v1 "github.com/Bruno-10/tasks/business/web/v1"
	"github.com/Bruno-10/tasks/business/web/v1/paging"
	"github.com/Bruno-10/tasks/foundation/web"
)

// Handlers manages the set of task endpoints.
type Handlers struct {
	task *task.Core
}

// New constructs a handlers for route access.
func New(task *task.Core) *Handlers {
	return &Handlers{
		task: task,
	}
}

// Create adds a new task to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewTask

	if err := web.Decode(r, &app); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	nc, err := toCoreNewTask(app)
	if err != nil {
		fmt.Println(err, "test")
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	tsk, err := h.task.Create(ctx, nc)
	if err != nil {
		return fmt.Errorf("create: tsk[%+v]: %w", tsk, err)
	}

	return web.Respond(ctx, w, toAppTask(tsk), http.StatusCreated)
}

// Query returns a list of tasks with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	tasks, err := h.task.Query(ctx)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppTask, len(tasks))
	for i, tsk := range tasks {
		items[i] = toAppTask(tsk)
	}

	total, err := h.task.Count(ctx)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, 1, 10000), http.StatusOK)
}
