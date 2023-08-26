// Package taskdb contains task related CRUD functionality.
package taskdb

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/data/order"
	database "github.com/Bruno-10/tasks/business/sys/database/pgx"
	"github.com/Bruno-10/tasks/business/sys/database/pgx/dbarray"
	"github.com/google/uuid"
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

// Update replaces a task document in the database.
func (s *Store) Update(ctx context.Context, tsk task.Task) error {
	const q = `
	UPDATE
		tasks
	SET 
		"label" = :label,
		"date_updated" = :date_updated
	WHERE
		task_id = :task_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBTask(tsk)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return task.ErrUniqueEmail
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a task from the database.
func (s *Store) Delete(ctx context.Context, tsk task.Task) error {
	data := struct {
		TaskID string `db:"task_id"`
	}{
		TaskID: tsk.ID.String(),
	}

	const q = `
	DELETE FROM
		tasks
	WHERE
		task_id = :task_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing tasks from the database.
func (s *Store) Query(ctx context.Context, filter task.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]task.Task, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		tasks`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbTsks []dbTask
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbTsks); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toCoreTaskSlice(dbTsks), nil
}

// Count returns the total number of tasks in the DB.
func (s *Store) Count(ctx context.Context, filter task.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
		tasks`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := database.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return count.Count, nil
}

// QueryByID gets the specified task from the database.
func (s *Store) QueryByID(ctx context.Context, taskID uuid.UUID) (task.Task, error) {
	data := struct {
		ID string `db:"task_id"`
	}{
		ID: taskID.String(),
	}

	const q = `
	SELECT
		*
	FROM
		tasks
	WHERE 
		task_id = :task_id`

	var dbTsk dbTask
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbTsk); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return task.Task{}, fmt.Errorf("namedquerystruct: %w", task.ErrNotFound)
		}
		return task.Task{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toCoreTask(dbTsk), nil
}

// QueryByIDs gets the specified tasks from the database.
func (s *Store) QueryByIDs(ctx context.Context, taskIDs []uuid.UUID) ([]task.Task, error) {
	ids := make([]string, len(taskIDs))
	for i, taskID := range taskIDs {
		ids[i] = taskID.String()
	}

	data := struct {
		TaskID interface {
			driver.Valuer
			sql.Scanner
		} `db:"task_id"`
	}{
		TaskID: dbarray.Array(ids),
	}

	const q = `
	SELECT
		*
	FROM
		tasks
	WHERE
		task_id = ANY(:task_id)`

	var tsks []dbTask
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &tsks); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return nil, task.ErrNotFound
		}
		return nil, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toCoreTaskSlice(tsks), nil
}

// QueryByLabel gets the specified task from the database by label.
func (s *Store) QueryByLabel(ctx context.Context, label string) (task.Task, error) {
	data := struct {
		Label string `db:"label"`
	}{
		Label: label,
	}

	const q = `
	SELECT
		*
	FROM
		tasks
	WHERE
		label = :label`

	var dbTsk dbTask
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbTsk); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return task.Task{}, fmt.Errorf("namedquerystruct: %w", task.ErrNotFound)
		}
		return task.Task{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toCoreTask(dbTsk), nil
}
