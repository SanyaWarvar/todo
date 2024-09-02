package repository

import (
	"fmt"
	"strings"

	"github.com/SanyaWarvar/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(
		`
		INSERT INTO %s (title, description, done) VALUES ($1, $2, $3) RETURNING id
		`,
		todoItemsTable,
	)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Done)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf(
		`
		INSERT INTO %s (list_id, item_id) VALUES ($1, $2) RETURNING id
		`,
		listsItemsTable,
	)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(
		`
		select 
		ti.id,
		ti.title, 
		ti.description,
		ti.done
		from %s ti
		inner join %s li on ti.id = li.item_id 
		inner join %s ul on ul.list_id = li.list_id 
		where ul.user_id = $1 and li.list_id = $2
		;
		`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	err := r.db.Select(&items, query, userId, listId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(
		`
		select 
		ti.id,
		ti.title, 
		ti.description,
		ti.done
		from %s ti
		inner join %s li on ti.id = li.item_id 
		inner join %s ul on ul.list_id = li.list_id 
		where ul.user_id = $1 and ti.id = $2
		`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	err := r.db.Get(&item, query, userId, itemId)

	return item, err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {

	query := fmt.Sprintf(
		`--sql
		DELETE FROM %s ti
		USING %s li, %s ul
		WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2
		`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	_, err := r.db.Exec(query, userId, itemId)
	return err

}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`
		UPDATE %s ti
		SET %s 
		FROM %s ul, %s li
		WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d
		`,
		todoItemsTable,
		setQuery,
		usersListsTable,
		listsItemsTable,
		argId,
		argId+1,
	)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err

}
