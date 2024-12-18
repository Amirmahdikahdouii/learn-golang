package services

import (
	"context"
	"errors"

	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/database"
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/models"
)

type TodoService interface {
	GetAllTodos() ([]models.Todo, error)
	CreateTodo(title string, completed bool) (models.Todo, error)
	UpdateTodo(id int, title string, completed bool) (models.Todo, error)
	DeleteTodo(id int) error
}

type todoService struct{}

func NewTodoService() TodoService {
	return &todoService{}
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM tasks ORDER BY id DESC")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *todoService) CreateTodo(title string, completed bool) (models.Todo, error) {
	var newTodo models.Todo
	err := database.DB.QueryRow(
		context.Background(),
		"INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id, title, completed",
		title, completed,
	).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Completed)

	if err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

func (s *todoService) UpdateTodo(id int, title string, completed bool) (models.Todo, error) {
	var updatedTodo models.Todo
	err := database.DB.QueryRow(
		context.Background(),
		"UPDATE tasks SET title = $1, completed = $2 WHERE id = $3 RETURNING id, title, completed",
		title, completed, id,
	).Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Completed)
	if err != nil {
		return models.Todo{}, err
	}
	return updatedTodo, nil
}

func (s *todoService) DeleteTodo(id int) error {
	commandTag, err := database.DB.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("todo task not found or deleted")
	}
	return nil
}
