// pkg/database/mongodb/repository/organization.go

package repository

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/yourusername/golang-api-project/pkg/database/mongodb/models"
)

// OrganizationRepository represents the repository for organization operations
type OrganizationRepository struct {
    collection *mongo.Collection
}

// NewOrganizationRepository creates a new instance of OrganizationRepository
func NewOrganizationRepository(db *mongo.Database) *OrganizationRepository {
    return &OrganizationRepository{
        collection: db.Collection("organizations"),
    }
}

// CreateOrganization creates a new organization
func (repo *OrganizationRepository) CreateOrganization(ctx context.Context, org *models.Organization) (*mongo.InsertOneResult, error) {
    return repo.collection.InsertOne(ctx, org)
}

// GetOrganizationByID retrieves an organization by ID
func (repo *OrganizationRepository) GetOrganizationByID(ctx context.Context, id primitive.ObjectID) (*models.Organization, error) {
    var org models.Organization
    err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&org)
    if err != nil {
        return nil, err
    }
    return &org, nil
}

// GetAllOrganizations retrieves all organizations
func (repo *OrganizationRepository) GetAllOrganizations(ctx context.Context) ([]*models.Organization, error) {
    cursor, err := repo.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var orgs []*models.Organization
    for cursor.Next(ctx) {
        var org models.Organization
        err := cursor.Decode(&org)
        if err != nil {
            return nil, err
        }
        orgs = append(orgs, &org)
    }
    return orgs, nil
}

// UpdateOrganization updates an existing organization
func (repo *OrganizationRepository) UpdateOrganization(ctx context.Context, id primitive.ObjectID, org *models.Organization) (*mongo.UpdateResult, error) {
    return repo.collection.ReplaceOne(ctx, bson.M{"_id": id}, org)
}

// DeleteOrganization deletes an organization by ID
func (repo *OrganizationRepository) DeleteOrganization(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    return repo.collection.DeleteOne(ctx, bson.M{"_id": id})
}
