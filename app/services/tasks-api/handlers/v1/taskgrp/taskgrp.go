// Package taskgrp maintains the group of handlers for task access.
package taskgrp

import (
	"context"
	"errors"
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
		return err
	}

	nc, err := toCoreNewTask(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	tsk, err := h.task.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, task.ErrUniqueEmail) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: tsk[%+v]: %w", tsk, err)
	}

	return web.Respond(ctx, w, toAppTask(tsk), http.StatusCreated)
}

// Update updates a task in the system.
// func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	var app AppUpdateTask
// 	if err := web.Decode(r, &app); err != nil {
// 		return err
// 	}

// 	taskID := auth.GetTaskID(ctx)

// 	tsk, err := h.task.QueryByID(ctx, taskID)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, task.ErrNotFound):
// 			return v1.NewRequestError(err, http.StatusNotFound)
// 		default:
// 			return fmt.Errorf("querybyid: taskID[%s]: %w", taskID, err)
// 		}
// 	}

// 	uu, err := toCoreUpdateTask(app)
// 	if err != nil {
// 		return v1.NewRequestError(err, http.StatusBadRequest)
// 	}

// 	tsk, err = h.task.Update(ctx, tsk, uu)
// 	if err != nil {
// 		return fmt.Errorf("update: taskID[%s] uu[%+v]: %w", taskID, uu, err)
// 	}

// 	return web.Respond(ctx, w, toAppTask(tsk), http.StatusOK)
// }

// Delete removes a task from the system.
// func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	taskID := auth.GetTaskID(ctx)

// 	tsk, err := h.task.QueryByID(ctx, taskID)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, task.ErrNotFound):
// 			return web.Respond(ctx, w, nil, http.StatusNoContent)
// 		default:
// 			return fmt.Errorf("querybyid: taskID[%s]: %w", taskID, err)
// 		}
// 	}

// 	if err := h.task.Delete(ctx, tsk); err != nil {
// 		return fmt.Errorf("delete: taskID[%s]: %w", taskID, err)
// 	}

// 	return web.Respond(ctx, w, nil, http.StatusNoContent)
// }

// Query returns a list of tasks with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := paging.ParseRequest(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	tasks, err := h.task.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppTask, len(tasks))
	for i, tsk := range tasks {
		items[i] = toAppTask(tsk)
	}

	total, err := h.task.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}

// QueryByID returns a task by its ID.
// func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	id := auth.GetTaskID(ctx)

// 	tsk, err := h.task.QueryByID(ctx, id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, task.ErrNotFound):
// 			return v1.NewRequestError(err, http.StatusNotFound)
// 		default:
// 			return fmt.Errorf("querybyid: id[%s]: %w", id, err)
// 		}
// 	}

// 	return web.Respond(ctx, w, toAppTask(tsk), http.StatusOK)
// }
