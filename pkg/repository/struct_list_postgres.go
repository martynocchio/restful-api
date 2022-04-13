package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restful_api "restful-api"
)

type StructListPostgres struct {
	db *sqlx.DB
}

func NewStructListPostgres(db *sqlx.DB) *StructListPostgres {
	return &StructListPostgres{db: db}
}

func (r *StructListPostgres) Create(userId int, list restful_api.StructList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", structListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *StructListPostgres) GetAll(userId int) ([]restful_api.StructList, error) {
	var lists []restful_api.StructList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		structListTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}