package todo

type User struct {
	Id       string `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
