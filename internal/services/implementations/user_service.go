package services

import (
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

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *UserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
