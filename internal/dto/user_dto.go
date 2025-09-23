package dto

import (
	"hoc-gin/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserResponse struct {
	UserID    int32     `json:"user_id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"create_at"`
}

// Response
func MapUserToDTO(user sqlc.User) *UserResponse {
	return &UserResponse{
		UserID:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
