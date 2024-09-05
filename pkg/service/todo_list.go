package service

import (
	"errors"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {

	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	_, err := s.repo.GetById(userId, listId)

	if err == nil {
		return s.repo.Delete(listId)
	}
	return err

}

func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if input.IsValid() {
		list, err := s.repo.GetById(userId, listId)
		if err != nil {
			return err
		}
		return s.repo.Update(list, input)
	}
	return errors.New("invalid input")
}
