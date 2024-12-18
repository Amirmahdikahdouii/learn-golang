package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/services"
	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetAllTodos(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) TodoHandler {
	return &todoHandler{service: service}
}

func (h *todoHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error"})
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *todoHandler) CreateTodo(c echo.Context) error {
	var request struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	newTodo, err := h.service.CreateTodo(request.Title, request.Completed)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error"})
	}
	return c.JSON(http.StatusCreated, newTodo)
}

func (h *todoHandler) UpdateTodo(c echo.Context) error {
	var request struct {
		ID        int    `param:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	updatedTodo, err := h.service.UpdateTodo(request.ID, request.Title, request.Completed)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error"})
	}
	return c.JSON(http.StatusOK, updatedTodo)
}

func (h *todoHandler) DeleteTodo(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	log.Printf("%v", ID)
	err := h.service.DeleteTodo(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "delete successfully"})
}
