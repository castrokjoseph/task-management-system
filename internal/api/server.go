package api

import (
	"fmt"
	"task-management-system/configs"
	"task-management-system/internal/api/handlers"
	"task-management-system/internal/api/middleware"
	"task-management-system/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config *configs.Config
}

func NewServer(
	authService *service.AuthService, 
	projectService *service.ProjectService, 
	taskService *service.TaskService,
) *Server {
	router := gin.Default()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Public routes
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(config))
	{
		// Task routes
		authorized.POST("/tasks", taskHandler.CreateTask)
		authorized.GET("/tasks", taskHandler.GetTasks)
		authorized.PUT("/tasks/:id", taskHandler.UpdateTask)
		authorized.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	return &Server{
		router: router,
		config: config,
	}
}

func (s *Server) Run(config *configs.Config) {
	port := config.Server.Port
	if port == 0 {
		port = 8080 // Default port
	}

	// Configure Gin mode based on environment
	gin.SetMode(gin.ReleaseMode)

	// Start server
	err := s.router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
