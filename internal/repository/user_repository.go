package repository

import (
	"database/sql"
	"fmt"
	"hoc-gin/internal/models"
	"log"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(DB *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: DB,
	}
}

func (ur *SQLUserRepository) Create(user *models.User) error {
	row := ur.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING user_id", user.Name, user.Email)
	if err := row.Scan(&user.Id); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (ur *SQLUserRepository) FindById(id int) {
	log.Println("Find by ID")
}
