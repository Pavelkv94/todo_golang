package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Pavelkv94/todo_golang/internal/core/domain"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{usersService: usersService}
}

func (h *UsersHTTPHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
