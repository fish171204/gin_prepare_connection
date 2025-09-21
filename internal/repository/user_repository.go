package repository

import (
	"database/sql"
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

func (ur *SQLUserRepository) Create(user *models.User) {
	log.Println("Create")
}

func (ur *SQLUserRepository) FindById(id int) {
	log.Println("Find by ID")
}
