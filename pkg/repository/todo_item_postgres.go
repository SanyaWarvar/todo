package repository

import (
	"errors"

	"github.com/SanyaWarvar/todo-app"
	"gorm.io/gorm"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewTodoItemPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {

	tx := r.db.Begin()

	result := tx.Table("todo_items").Create(&item)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	listItems := todo.ListsItem{
		ListId: listId,
		ItemId: item.Id,
	}

	result = tx.Table("list_items").Create(&listItems)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	result = tx.Commit()

	return item.Id, result.Error
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	result := r.db.Table("todo_items ti").
		Select("ti.id", "ti.title", "ti.description", "ti.done").
		Joins("inner join list_items li on ti.id = li.item_id").
		Joins("inner join users_lists ul on ul.list_id = li.list_id").
		Where("ul.user_id = ? and li.list_id = ?", userId, listId).
		Scan(&items)

	return items, result.Error
}
func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	result := r.db.Table("todo_items ti").
		Select("ti.id, ti.title, ti.description, ti.done").
		Joins("inner join list_items li on ti.id = li.item_id").
		Joins("inner join users_lists ul on ul.list_id = li.list_id").
		Where("ul.user_id = ? and ti.id = ?", userId, itemId).
		Find(&item)
	if item.Id == 0 { //запись не найдена
		return item, errors.New("not found")
	}
	return item, result.Error
}

func (r *TodoItemPostgres) Delete(item todo.TodoItem) error {

	result := r.db.Table("todo_items").Delete(&item)
	return result.Error

}

func (r *TodoItemPostgres) Update(item todo.TodoItem, input todo.UpdateItemInput) error {

	if input.Title != nil {
		item.Title = *input.Title
	}

	if input.Description != nil {
		item.Description = *input.Description
	}

	if input.Done != nil {
		item.Done = *input.Done
	}

	result := r.db.Save(&item)
	return result.Error
}
