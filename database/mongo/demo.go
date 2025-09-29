package mongo

import (
	"templates/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type demoRepositoryIml struct {
	collection *mongo.Collection
}

func NewDemoRepository() database.DemoRepository {
	return &demoRepositoryIml{}
}
