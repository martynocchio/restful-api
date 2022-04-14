package repository

import (
	"github.com/jmoiron/sqlx"
	restful_api "restful-api"
)

type Authorization interface {
	CreateUser(user restful_api.User) (int, error)
	GetUser(username, password string) (restful_api.User, error)
}

type StructList interface {
	Create(userId int, list restful_api.StructList) (int, error)
	GetAll(userId int) ([]restful_api.StructList, error)
	GetById(userId, listId int) (restful_api.StructList, error)
	Update(userId, listId int, input restful_api.UpdateListInput) error
	Delete(userId, listId int) error
}

type StructTask interface {
	Create(listId int, task restful_api.StructTask) (int, error)
	GetAll(userId, listId int) ([]restful_api.StructTask, error)
}

type StructSubtask interface {
}

type Repository struct {
	Authorization
	StructList
	StructTask
	StructSubtask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		StructList:    NewStructListPostgres(db),
		StructTask:    NewStructTaskPostgres(db),
	}
}
