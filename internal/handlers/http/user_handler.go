package http

import (
	"strconv"

	gin "github.com/gin-gonic/gin"
	requests "minh.com/go-rest-gin-3/internal/handlers/dtos/requests"
	factories "minh.com/go-rest-gin-3/internal/handlers/factories"
	mapper "minh.com/go-rest-gin-3/internal/handlers/mappers"
	services "minh.com/go-rest-gin-3/internal/services/interfaces"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(
	userService services.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the system by providing the required details in the request body.
// @Tags users
// @Accept json
// @Produce json
// @Param body body requests.CreateUserRequest true "User creation payload"
// @Success 201 {object} user.UserResponse "User created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		factories.BadRequest(c, err)
		return
	}

	userEntity := mapper.CreateRequestToEntity(&req)

	newUser, err := uh.userService.CreateUser(userEntity)
	if err != nil {
		factories.InternalServerError(c, err)
		return
	}

	factories.Created(
		c,
		"New user created successfully",
		mapper.EntityToResponse(newUser),
	)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve a single user by their ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} user.UserResponse "User retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users/{id} [get]
func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		factories.BadRequest(c, "Invalid user ID")
		return
	}

	user, err := uh.userService.GetUserByID(id)
	if err != nil {
		factories.NotFound(c, "User")
		return
	}

	factories.SuccessSingle(
		c,
		"User retrieved successfully",
		mapper.EntityToResponse(user),
	)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} user.UserResponse "Users retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [get]
func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.userService.GetAllUsers()
	if err != nil {
		factories.InternalServerError(c, err)
		return
	}

	factories.SuccessList(c, "Users retrieved successfully", mapper.EntitiesToResponse(users))
}

// 	UpdateUser(id int, user *models.User) (*models.User, error)
// 	DeleteUser(id int) error
