package restful_api

type structList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersLists struct {
	Id     int
	UserId int
	ListId int
}

type structTask struct {
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

type structSubtask struct {
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
