package repositories

import (
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"minh.com/go-rest-gin-3/internal/models"
)

type UserRepository struct {
	collections *mongo.Collection
}

// Constructor
func NewUserRepository(collections *mongo.Collection) *UserRepository {
	return &UserRepository{
		collections: collections,
	}
}

/*
(r *UserRepository) - Receiver
r: Tên biến receiver (có thể đặt tên bất kỳ, thường là chữ cái đầu của struct)
*UserRepository: Kiểu receiver - là pointer đến struct UserRepository
Ý nghĩa: Method này "thuộc về" struct UserRepository
*/
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	res, err := r.collections.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collections.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := r.collections.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, updated_user *models.User) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updated_user.UpdatedAt = time.Now()
	_, err = r.collections.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updated_user})
	if err != nil {
		return nil, err
	}

	return updated_user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collections.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
