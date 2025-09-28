package services

import (
	"context"

	"minh.com/go-rest-gin-3/internal/models"
	repositories "minh.com/go-rest-gin-3/internal/repositories/interfaces"
)

type UserService struct {
	// inject repository
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepository.GetUserByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepository.GetAllUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepository.DeleteUser(ctx, id)
}
