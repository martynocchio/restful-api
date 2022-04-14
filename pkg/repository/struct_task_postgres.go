package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restful_api "restful-api"
)

type StructTaskPostgres struct {
	db *sqlx.DB
}

func NewStructTaskPostgres(db *sqlx.DB) *StructTaskPostgres {
	return &StructTaskPostgres{db: db}
}

func (r *StructTaskPostgres) Create(listId int, task restful_api.StructTask) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		structTaskTable)

	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)
	err = row.Scan(&taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListTasksQuery := fmt.Sprintf("INSERT INTO %s (list_id, task_id) VALUES ($1, $2)",
		listsTasksTable)
	_, err = tx.Exec(createListTasksQuery, listId, taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskId, tx.Commit()
}

func (r *StructTaskPostgres) GetAll(userId, listId int) ([]restful_api.StructTask, error) {
	var tasks []restful_api.StructTask
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti
								 INNER JOIN %s li on li.task_id = ti.id
								 INNER JOIN %s ul on ul.list_id = li.list_id
								 WHERE li.list_id = $1 AND ul.user_id = $2`,
		structTaskTable, listsTasksTable, usersListsTable)

	if err := r.db.Select(&tasks, query, listId, userId); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *StructTaskPostgres) GetById(userId, taskId int) (restful_api.StructTask, error) {
	var task restful_api.StructTask

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti
								INNER JOIN %s li on li.task_id = ti.id
								INNER JOIN %s ul on ul.list_id = li.list_id
								WHERE ti.id = $1 AND ul.user_id = $2`,
		structTaskTable, listsTasksTable, usersListsTable)
	if err := r.db.Get(&task, query, taskId, userId); err != nil {
		return task, err
	}

	return task, nil
}
