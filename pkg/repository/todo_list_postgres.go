package repository

import (
	"github.com/SanyaWarvar/todo-app"
	"gorm.io/gorm"
)

type TodoListPostgres struct {
	db *gorm.DB
}

func NewTodoListPostgres(db *gorm.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx := r.db.Begin()

	result := tx.Create(&list)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	result = tx.Create(&todo.UsersList{
		UserId: userId,
		ListId: list.Id,
	})
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	result = tx.Commit()

	return list.Id, result.Error
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	result := r.db.Table(todoListsTable).
		Select("todo_lists.id", "todo_lists.title", "todo_lists.description").
		Joins("inner join users_lists ul ON todo_lists.id = ul.list_id").
		Where("ul.user_id = ?", userId).
		Scan(&lists)

	return lists, result.Error
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	result := r.db.Select("todo_lists.id", "todo_lists.title", "todo_lists.description").
		Table(todoListsTable).
		Joins("inner join users_lists on users_lists.list_id = todo_lists.id").
		Where("users_lists.user_id = ? AND users_lists.list_id = ?", userId, listId).
		Find(&list)

	return list, result.Error
}

func (r *TodoListPostgres) Delete(listId int) error {
	result := r.db.Table(todoListsTable).Delete(&todo.TodoList{Id: listId})
	return result.Error
}

func (r *TodoListPostgres) Update(list todo.TodoList, input todo.UpdateListInput) error {
	if input.Title != nil {
		list.Title = *input.Title
	}
	if input.Description != nil {
		list.Description = *input.Description
	}
	result := r.db.Save(list)
	return result.Error
}
