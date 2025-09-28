package user

import "minh.com/go-rest-gin-3/internal/handlers/dtos/responses"

type UserResponse struct {
	responses.BaseDataResponse
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}
