package taskgrp

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (task.QueryFilter, error) {
	values := r.URL.Query()

	var filter task.QueryFilter

	if taskID := values.Get("task_id"); taskID != "" {
		id, err := uuid.Parse(taskID)
		if err != nil {
			return task.QueryFilter{}, validate.NewFieldsError("task_id", err)
		}
		filter.WithTaskID(id)
	}

	if email := values.Get("email"); email != "" {
		addr, err := mail.ParseAddress(email)
		if err != nil {
			return task.QueryFilter{}, validate.NewFieldsError("email", err)
		}
		filter.WithEmail(*addr)
	}

	if createdDate := values.Get("start_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return task.QueryFilter{}, validate.NewFieldsError("start_created_date", err)
		}
		filter.WithStartDateCreated(t)
	}

	if createdDate := values.Get("end_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return task.QueryFilter{}, validate.NewFieldsError("end_created_date", err)
		}
		filter.WithEndCreatedDate(t)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if err := filter.Validate(); err != nil {
		return task.QueryFilter{}, err
	}

	return filter, nil
}
