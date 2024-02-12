// pkg/database/mongodb/models/organization.go

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Organization represents an organization entity
type Organization struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Description string             `bson:"description"`
    Members     []Member           `bson:"members"`
}

// Member represents a member of an organization
type Member struct {
    Name        string `bson:"name"`
    Email       string `bson:"email"`
    AccessLevel string `bson:"access_level"`
}
