package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
