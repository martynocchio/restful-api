package restful_api

import "errors"

type StructList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersLists struct {
	Id     int
	UserId int
	ListId int
}

type StructTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsTask struct {
	Id     int
	ListId int
	TaskId int
}

type StructSubtask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TasksSubtask struct {
	Id        int
	TaskId    int
	SubtaskId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("UPDATE STRUCTURE HAS NO VALUES")
	}

	return nil
}
