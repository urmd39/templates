package mongo

import (
	"templates/database"
	"templates/infrastructure"
	"templates/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type tokenRepositoryIml struct {
	collection *mongo.Collection
}

func NewTokenRepository(db *infrastructure.MongoDBStore, collection string) database.TokenRepository {
	return &tokenRepositoryIml{
		collection: db.Db.Collection(collection),
	}
}

func (r *tokenRepositoryIml) UpsertToken(token *model.Token) (err error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "collection_name", Value: token.CollectionName}}
	update := bson.D{{Key: "$set", Value: token}}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (r *tokenRepositoryIml) GetTokenByNameCollection(name string) (token *model.Token, err error) {
	filter := bson.D{{Key: "collection_name", Value: name}}

	token = &model.Token{}
	err = r.collection.FindOne(context.TODO(), filter).Decode(token)
	return
}

func (r *tokenRepositoryIml) DeleteToken(collection string) (err error) {
	filter := bson.M{"collection_name": collection}
	_, err = r.collection.DeleteOne(context.TODO(), filter)
	return
}
