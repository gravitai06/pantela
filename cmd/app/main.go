package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pantela/internal/database"
	"pantela/internal/handlers"
	"pantela/internal/taskServise"
	"pantela/internal/web/tasks"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskServise.Task{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	repo := taskServise.NewRepository(database.DB)
	service := taskServise.NewService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil) // Второй аргумент — это опции, можно оставить nil
	tasks.RegisterHandlers(e, strictHandler)

	log.Println("Server started on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
