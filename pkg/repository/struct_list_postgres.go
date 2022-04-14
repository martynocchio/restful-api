package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	restful_api "restful-api"
	"strings"
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

func (r *StructListPostgres) GetById(userId, listId int) (restful_api.StructList, error) {
	var list restful_api.StructList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul
					on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		structListTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *StructListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 
						AND ul.list_id = $2`, structListTable, usersListsTable)

	_, err := r.db.Exec(query, userId, listId)
	return err
}

func (r *StructListPostgres) Update(userId, listId int, input restful_api.UpdateListInput) error {
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul
								WHERE tl.id = ul.list_id AND
								ul.list_id=$%d AND ul.user_id=$%d`,
		structListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
