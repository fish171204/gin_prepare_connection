package repository

import (
	"hoc-gin/internal/db/sqlc"
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

func (ur *SQLUserRepository) Create(input sqlc.CreateUserParams) {
	log.Println("Create")
}

func (ur *SQLUserRepository) FindById(id int) {
	log.Println("Find by ID")
}
