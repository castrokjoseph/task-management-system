package service

import (
	"task-management-system/internal/models"
	"task-management-system/internal/repository"

	"github.com/google/uuid"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(userID uuid.UUID, taskCreate *models.TaskCreate) (*models.Task, error) {
	task := &models.Task{
		Title:       taskCreate.Title,
		Description: taskCreate.Description,
		Status:      taskCreate.Status,
		Priority:    taskCreate.Priority,
		UserID:      userID,
		ProjectID:   taskCreate.ProjectID,
		Deadline:    taskCreate.Deadline,
	}

	err := s.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetUserTasks(userID uuid.UUID) ([]models.Task, error) {
	return s.taskRepo.FindByUserID(userID)
}

func (s *TaskService) UpdateTask(userID uuid.UUID, task *models.Task) error {
	// Ensure the task belongs to the user
	task.UserID = userID
	return s.taskRepo.Update(task)
}

func (s *TaskService) DeleteTask(taskID, userID uuid.UUID) error {
	return s.taskRepo.Delete(taskID, userID)
}