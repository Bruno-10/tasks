package taskdb

import (
	"time"

	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/google/uuid"
)

// dbTask represent the structure we need for moving data
// between the app and the database.
type dbTask struct {
	ID          uuid.UUID `db:"task_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Type        string    `db:"type"`
	Label       string    `db:"label"`
	DueDate     time.Time `db:"due_date"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBTask(tsk task.Task) dbTask {

	return dbTask{
		ID:          tsk.ID,
		Name:        tsk.Name,
		Description: tsk.Description,
		Type:        tsk.Type,
		Label:       tsk.Label,
		DueDate:     tsk.DueDate.UTC(),
		DateCreated: tsk.DateCreated.UTC(),
		DateUpdated: tsk.DateUpdated.UTC(),
	}
}

func toCoreTask(dbTsk dbTask) task.Task {

	tsk := task.Task{
		ID:          dbTsk.ID,
		Name:        dbTsk.Name,
		Description: dbTsk.Description,
		Type:        dbTsk.Type,
		Label:       dbTsk.Label,
		DueDate:     dbTsk.DueDate.In(time.Local),
		DateCreated: dbTsk.DateCreated.In(time.Local),
		DateUpdated: dbTsk.DateUpdated.In(time.Local),
	}

	return tsk
}

func toCoreTaskSlice(dbTasks []dbTask) []task.Task {
	tsks := make([]task.Task, len(dbTasks))
	for i, dbTsk := range dbTasks {
		tsks[i] = toCoreTask(dbTsk)
	}
	return tsks
}
