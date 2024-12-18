package routes

import (
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/handlers"
	"github.com/labstack/echo/v4"
)

type TodoRoutes interface {
	RegisterRoutes(e *echo.Echo)
}

type todoRoutes struct {
	handler handlers.TodoHandler
}

func NewTodoRoutes(handler handlers.TodoHandler) TodoRoutes {
	return &todoRoutes{handler: handler}
}

func (r *todoRoutes) RegisterRoutes(e *echo.Echo) {
	e.GET("/todo/all/", r.handler.GetAllTodos)
	e.POST("/todo/create/", r.handler.CreateTodo)
	e.PUT("/todo/update/:id", r.handler.UpdateTodo)
	e.DELETE("/todo/delete/:id", r.handler.DeleteTodo)
}
