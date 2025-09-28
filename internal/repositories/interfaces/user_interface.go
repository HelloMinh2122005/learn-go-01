package repositories

import (
	"context"

	"minh.com/go-rest-gin-3/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}
