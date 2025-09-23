package handlers

import (
	"hoc-gin/internal/db/sqlc"
	"hoc-gin/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invali user ID"})
		return
	}

	uh.repo.FindById(id)

	ctx.JSON(http.StatusOK, gin.H{"data": "Get user by uuid"})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uh.repo.Create(input)

	ctx.JSON(http.StatusCreated, gin.H{"data": "Create user"})
}
