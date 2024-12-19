package main

import (
	"log"
	"task-management-system/configs"
	"task-management-system/internal/api"
	"task-management-system/internal/repository"
	"task-management-system/internal/service"
)

func main() {
	config := configs.LoadConfig()

	// Initialize database connection
	db, err := repository.NewPostgresConnection(config)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, config)
	projectService := service.NewProjectService(projectRepo)
	taskService := service.NewTaskService(taskRepo)

	// Setup and run server
	server := api.NewServer(authService, projectService, taskService)
	server.Run(config)
}