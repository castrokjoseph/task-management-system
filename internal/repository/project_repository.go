package repository

import (
	"task-management-system/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	query := `INSERT INTO projects (name, description, user_id) 
              VALUES (:name, :description, :user_id) 
              RETURNING id, created_at`
	
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"name":        project.Name,
		"description": project.Description,
		"user_id":     project.UserID,
	})
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&project.ID, &project.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ProjectRepository) FindByUserID(userID uuid.UUID) ([]models.Project, error) {
	var projects []models.Project
	query := `SELECT * FROM projects WHERE user_id = $1 ORDER BY created_at DESC`
	
	err := r.db.Select(&projects, query, userID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}