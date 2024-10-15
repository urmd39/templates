package database

import (
	"encoding/json"
	"templates/database"
	"templates/infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IWorkerDatabase interface {
	Listen()
	CallBack(document *Document)
}

type BaseWorkerDatabase struct {
	collection      *mongo.Collection
	eventPusher     infrastructure.EventPublisher
	tokenRepository database.TokenRepository
}

const (
	Create  = "insert"
	Update  = "update"
	Replace = "replace"
	Delete  = "delete"
)

type Document struct {
	ID struct {
		Data string `bson:"_data"`
	} `bson:"_id"`
	//Action string
	OperationType string `bson:"operationType"`
	FullDocument  bson.M `json:"fullDocument"`
	DocumentID    struct {
		ID string `bson:"_id"`
	} `bson:"documentKey"`
	ClusterTime primitive.Timestamp `json:"clusterTime"`
}

func convertData(input interface{}, output interface{}) error {
	dataByte, err := json.Marshal(input)
	if err != nil {
		infrastructure.ErrLog.Print(err)
		return err
	}

	if err = json.Unmarshal(dataByte, &output); err != nil {
		infrastructure.ErrLog.Print(err)
		return err
	}

	return err
}
