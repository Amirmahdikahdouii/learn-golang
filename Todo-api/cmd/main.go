package main

import (
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/database"
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/handlers"
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/routes"
	"github.com/amirmahdikahdouii/learn-golang/Todo-api/internal/services"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase("postgres://admin:123456@localhost:5433/todo_db")
	
	e := echo.New()
	todoService := services.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)
	todoRoutes := routes.NewTodoRoutes(todoHandler)

	todoRoutes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
