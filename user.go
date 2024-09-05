package todo

type User struct {
	Id            int    `json:"-" db:"id"`
	Username      string `json:"username" db:"username" binding:"required"`
	Password_hash string `json:"password" db:"password_hash" binding:"required"`
}
