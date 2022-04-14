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

	var itemId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		structTaskTable)

	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListTasksQuery := fmt.Sprintf("INSERT INTO %s (list_id, task_id) VALUES ($1, $2)",
		listsTasksTable)
	_, err = tx.Exec(createListTasksQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}
