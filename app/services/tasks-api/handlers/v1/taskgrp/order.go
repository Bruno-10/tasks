package taskgrp

import (
	"errors"
	"net/http"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/data/order"
	"github.com/Bruno-10/tasks/business/sys/validate"
)

var orderByFields = map[string]struct{}{
	task.OrderByID:      {},
	task.OrderByName:    {},
	task.OrderByEmail:   {},
	task.OrderByRoles:   {},
	task.OrderByEnabled: {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, task.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
