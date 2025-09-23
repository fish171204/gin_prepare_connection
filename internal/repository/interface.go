package repository

import "hoc-gin/internal/db/sqlc"

type UserRepository interface {
	Create(input sqlc.CreateUserParams)
	FindById(id int)
}
