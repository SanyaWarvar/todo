package repository

import (
	"github.com/SanyaWarvar/todo-app"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	user := todo.User{
		Username:      username,
		Password_hash: password,
	}
	result := r.db.Where(&user).First(&user)

	return user, result.Error
}
