package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Get user by uuid"})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Create user"})
}
