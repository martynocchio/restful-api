package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restful_api "restful-api"
	"strings"
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

func (r *StructTaskPostgres) Delete(userId, taskId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
								WHERE ti.id = li.task_id AND li.list_id = ul.list_id
								AND ul.user_id = $1 AND ti.id = $2`,
		structTaskTable, listsTasksTable, usersListsTable)
	_, err := r.db.Exec(query, userId, taskId)
	return err
}

func (r *StructTaskPostgres) Update(userId, taskId int, input restful_api.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
								WHERE ti.id = li.task_id AND
								li.list_id = ul.list_id AND
								ul.user_id=$%d AND ti.id=$%d`,
		structTaskTable, setQuery, listsTasksTable, usersListsTable, argId, argId+1)
	args = append(args, userId, taskId)

	_, err := r.db.Exec(query, args...)
	return err
}
