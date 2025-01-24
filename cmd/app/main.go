package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pantela/internal/database"
	"pantela/internal/handlers"
	"pantela/internal/taskServise"
	userService "pantela/internal/userServise"
	"pantela/internal/web/tasks"
	"pantela/internal/web/users"
)

func main() {
	database.InitDB()

	taskRepo := taskServise.NewRepository(database.DB)
	if err := taskRepo.Migrate(); err != nil {
		log.Fatalf("Failed to auto-migrate tasks database: %v", err)
	}
	taskService := taskServise.NewService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewRepository(database.DB)
	if err := userRepo.Migrate(); err != nil {
		log.Fatalf("Failed to auto-migrate users database: %v", err)
	}
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	log.Println("Server started on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
