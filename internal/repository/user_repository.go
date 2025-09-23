package repository

import (
	"context"
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

func (ur *SQLUserRepository) Create(ctx context.Context, input sqlc.CreateUserParams) {
	ur.db.CreateUser(ctx, input)

	log.Println("Create")
}

func (ur *SQLUserRepository) FindById(id int) {
	log.Println("Find by ID")
}
