package mappers

import (
	requests "minh.com/go-rest-gin-3/internal/handlers/dtos/requests"
	base "minh.com/go-rest-gin-3/internal/handlers/dtos/responses"
	responses "minh.com/go-rest-gin-3/internal/handlers/dtos/responses/user"
	models "minh.com/go-rest-gin-3/internal/models"
)

func CreateRequestToEntity(req *requests.CreateUserRequest) *models.User {
	// TODO: Validate fields and hash pashword
	return &models.User{
		ID:        0,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}

func UpdateRequestToEntity(req *requests.UpdateUserRequest) *models.User {
	return &models.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
}

func EntityToResponse(user *models.User) *responses.UserResponse {
	return &responses.UserResponse{
		BaseDataResponse: base.BaseDataResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func EntitiesToResponse(users []*models.User) []responses.UserResponse {
	var usersResponse []responses.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, *EntityToResponse(user))
	}
	return usersResponse
}
