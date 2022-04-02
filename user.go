package restful_api

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"firstname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
