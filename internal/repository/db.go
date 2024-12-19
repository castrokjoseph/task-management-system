package repository

import (
	"fmt"
	"task-management-system/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(config *configs.Config) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}
