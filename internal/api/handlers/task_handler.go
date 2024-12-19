package handlers

import (
	"net/http"
	"task-management-system/internal/models"
	"task-management-system/internal/service"
	"task-management-system/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	
	var taskCreate models.TaskCreate
	if err := c.ShouldBindJSON(&taskCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if err := utils.Validate(taskCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.taskService.CreateTask(userID, &taskCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)

	tasks, err := h.taskService.GetUserTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	
	taskIDStr := c.Param("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = taskID

	err = h.taskService.UpdateTask(userID, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	
	taskIDStr := c.Param("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = h.taskService.DeleteTask(taskID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}