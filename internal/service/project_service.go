package service

import (
	"task-management-system/internal/models"
	"task-management-system/internal/repository"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
}

func NewProjectService(projectRepo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

func (s *ProjectService) CreateProject(userID uuid.UUID, name, description string) (*models.Project, error) {
	project := &models.Project{
		Name:        name,
		Description: description,
		UserID:      userID,
	}

	err := s.projectRepo.Create(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetUserProjects(userID uuid.UUID) ([]models.Project, error) {
	return s.projectRepo.FindByUserID(userID)
}