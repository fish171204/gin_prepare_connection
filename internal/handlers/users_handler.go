package handlers

import (
	"hoc-gin/internal/db/sqlc"
	"hoc-gin/internal/dto"
	"hoc-gin/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")

	userUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invali user ID"})
		return
	}

	user, err := uh.repo.FindByUuid(ctx, userUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.MapUserToDTO(user)

	ctx.JSON(http.StatusOK, gin.H{"data": userDTO})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.repo.Create(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.MapUserToDTO(user)

	ctx.JSON(http.StatusCreated, gin.H{"data": userDTO})
}
