package repository

import (
	"hoc-gin/internal/db/sqlc"
	"hoc-gin/internal/models"
	"log"
)

type SQLUserRepository struct {
	db *sqlc.Queries
}

func NewSQLUserRepository(DB *sqlc.Queries) UserRepository {
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
