package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type ProjectCreate struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}