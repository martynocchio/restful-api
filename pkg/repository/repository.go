package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
