package repository

import (
	"errors"
	"task-management-system/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	query := `INSERT INTO tasks 
              (title, description, status, priority, user_id, project_id, deadline) 
              VALUES (:title, :description, :status, :priority, :user_id, :project_id, :deadline) 
              RETURNING id, created_at, updated_at`
	
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"priority":    task.Priority,
		"user_id":     task.UserID,
		"project_id":  task.ProjectID,
		"deadline":    task.Deadline,
	})
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *TaskRepository) FindByUserID(userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	query := `SELECT * FROM tasks WHERE user_id = $1 ORDER BY created_at DESC`
	
	err := r.db.Select(&tasks, query, userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	query := `UPDATE tasks 
              SET title = :title, 
                  description = :description, 
                  status = :status, 
                  priority = :priority, 
                  project_id = :project_id, 
                  deadline = :deadline,
                  updated_at = CURRENT_TIMESTAMP
              WHERE id = :id`
	
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":          task.ID,
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"priority":    task.Priority,
		"project_id":  task.ProjectID,
		"deadline":    task.Deadline,
	})

	return err
}

func (r *TaskRepository) Delete(taskID, userID uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1 AND user_id = $2`
	
	result, err := r.db.Exec(query, taskID, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("task not found or unauthorized")
	}

	return nil
}