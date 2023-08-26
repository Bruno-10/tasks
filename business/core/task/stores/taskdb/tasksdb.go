// Package taskdb contains task related CRUD functionality.
package taskdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Bruno-10/tasks/business/core/task"
	database "github.com/Bruno-10/tasks/business/sys/database/pgx"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for task database access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new task into the database.
func (s *Store) Create(ctx context.Context, tsk task.Task) error {
	s.log.Infoln("I'm here")
	const q = `
	INSERT INTO tasks
		(task_id, name, description, type, due_date, label, date_created, date_updated)
	VALUES
		(:task_id, :name, :description, :type, :due_date, :label, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBTask(tsk)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", task.ErrUniqueEmail)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing tasks from the database.
func (s *Store) Query(ctx context.Context) ([]task.Task, error) {
	data := map[string]interface{}{
		"offset":        0,
		"rows_per_page": 10000,
	}

	const q = `
	SELECT
		*
	FROM
		tasks`

	var dbTsks []dbTask
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &dbTsks); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toCoreTaskSlice(dbTsks), nil
}

// Count returns the total number of tasks in the DB.
func (s *Store) Count(ctx context.Context) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
		tasks`

	var count struct {
		Count int `db:"count"`
	}
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &count); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return count.Count, nil
}
