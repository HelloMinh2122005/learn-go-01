package repositories

import (
	"minh.com/go-rest-gin-3/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
