package service

import (
	"errors"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	item, err := s.repo.GetById(userId, itemId)
	if err == nil {
		return s.repo.Delete(item)
	}
	return err
}

func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	if input.IsValid() {
		item, err := s.repo.GetById(userId, itemId)
		if err != nil {
			return err
		}
		return s.repo.Update(item, input)
	}
	return errors.New("invalid input")
}
