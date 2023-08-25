package task

import "github.com/Bruno-10/tasks/business/data/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByID      = "taskid"
	OrderByName    = "name"
	OrderByEmail   = "email"
	OrderByRoles   = "roles"
	OrderByEnabled = "enabled"
)
