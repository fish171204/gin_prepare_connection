package repository

import (
	"context"
	"hoc-gin/internal/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, input sqlc.CreateUserParams)
	FindById(id int)
}
