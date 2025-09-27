package user

import "minh.com/go-rest-gin-3/internal/handlers/dtos/responses"

type UserResponse struct {
	responses.BaseDataResponse
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
