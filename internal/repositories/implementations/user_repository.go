package repositories

import (
	"errors"
	"time"

	"minh.com/go-rest-gin-3/internal/models"
)

type UserRepository struct {
	users  []*models.User
	nextID int
}

// Constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make([]*models.User, 0),
		nextID: 1,
	}
}

func (r *UserRepository) SeedDummyData() error {
	dummyUsers := []*models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Alice",
			LastName:  "Johnson",
			Email:     "alice.johnson@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Bob",
			LastName:  "Williams",
			Email:     "bob.williams@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Charlie",
			LastName:  "Brown",
			Email:     "charlie.brown@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Diana",
			LastName:  "Davis",
			Email:     "diana.davis@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Eve",
			LastName:  "Miller",
			Email:     "eve.miller@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Frank",
			LastName:  "Wilson",
			Email:     "frank.wilson@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Grace",
			LastName:  "Moore",
			Email:     "grace.moore@example.com",
			Password:  "password123",
		},
		{
			FirstName: "Henry",
			LastName:  "Taylor",
			Email:     "henry.taylor@example.com",
			Password:  "password123",
		},
	}

	for _, user := range dummyUsers {
		if _, err := r.CreateUser(user); err != nil {
			return err // Return error if any creation fails
		}
	}

	return nil
}

/*
(r *UserRepository) - Receiver
r: Tên biến receiver (có thể đặt tên bất kỳ, thường là chữ cái đầu của struct)
*UserRepository: Kiểu receiver - là pointer đến struct UserRepository
Ý nghĩa: Method này "thuộc về" struct UserRepository
*/
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user can not be nil")
	}

	user.ID = r.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.nextID++
	r.users = append(r.users, user)

	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("no user found")
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	return r.users, nil
}

func (r *UserRepository) UpdateUser(id int, updated_user *models.User) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			user.Email = updated_user.Email
			user.FirstName = updated_user.FirstName
			user.LastName = updated_user.LastName
			user.UpdatedAt = time.Now()
			return user, nil
		}
	}

	return nil, errors.New("no user found")
}

func (r *UserRepository) DeleteUser(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}

	return errors.New("no user found")
}
