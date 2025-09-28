package mappers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	requests "minh.com/go-rest-gin-3/internal/handlers/dtos/requests"
	base "minh.com/go-rest-gin-3/internal/handlers/dtos/responses"
	responses "minh.com/go-rest-gin-3/internal/handlers/dtos/responses/user"
	models "minh.com/go-rest-gin-3/internal/models"
)

func CreateRequestToEntity(req *requests.CreateUserRequest) *models.User {
	// TODO: Validate fields and hash pashword
	return &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}

func UpdateRequestToEntity(req *requests.UpdateUserRequest) *models.User {
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil
	}
	return &models.User{
		ID:        objID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
}

func EntityToResponse(user *models.User) *responses.UserResponse {
	return &responses.UserResponse{
		BaseDataResponse: base.BaseDataResponse{
			ID:        user.ID.Hex(),
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
