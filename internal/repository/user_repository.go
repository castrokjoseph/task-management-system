package repository

import (
	"task-management-system/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash) 
              VALUES (:username, :email, :password_hash) 
              RETURNING id, created_at`
	
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"username":       user.Username,
		"email":          user.Email,
		"password_hash":  user.Password,
	})
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash, created_at 
              FROM users WHERE email = $1`
	
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}