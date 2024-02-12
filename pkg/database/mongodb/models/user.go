// pkg/database/mongodb/models/user.go

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user entity
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Name     string             `bson:"name"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
}
