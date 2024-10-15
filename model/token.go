package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	CollectionName string              `bson:"collection_name"`
	Token          string              `bson:"token"`
	ClusterTime    primitive.Timestamp `bson:"cluster_time"`
}
