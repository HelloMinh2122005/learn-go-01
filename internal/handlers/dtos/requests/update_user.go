package requests

type UpdateUserRequest struct {
	ID        int    `json:"ID" binding:"required"`
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	Email     string `json:"Email" binding:"required,email"`
}
