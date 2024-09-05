package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "list_items"
)

func NewPostgresDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=qwerty port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
