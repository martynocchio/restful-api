package repository

import (
	"github.com/jmoiron/sqlx"
	restful_api "restful-api"
)

type Authorization interface {
	CreateUser(user restful_api.User) (int, error)
}

type List interface {
}

type Task interface {
}

type Subtask interface {
}

type Repository struct {
	Authorization
	List
	Task
	Subtask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
