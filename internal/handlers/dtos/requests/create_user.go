package requests

type CreateUserRequest struct {
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	Email     string `json:"Email" binding:"required,email"`
	Password  string `json:"Password" binding:"required,min=6"`
}
