// pkg/database/mongodb/repository/user.go

package repository

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/yourusername/golang-api-project/pkg/database/mongodb/models"
)

// UserRepository represents the repository for user operations
type UserRepository struct {
    collection *mongo.Collection
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        collection: db.Collection("users"),
    }
}

// CreateUser creates a new user
func (repo *UserRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
    return repo.collection.InsertOne(ctx, user)
}

// GetUserByEmail retrieves a user by email
func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
