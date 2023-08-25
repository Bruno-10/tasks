package taskdb

import (
	"fmt"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/data/order"
)

var orderByFields = map[string]string{
	task.OrderByID:      "task_id",
	task.OrderByName:    "name",
	task.OrderByEmail:   "email",
	task.OrderByRoles:   "roles",
	task.OrderByEnabled: "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
