package repositories

import (
	"github.com/globalsign/mgo"
)

type MongoDBRepository struct {
	CollectionName string
	Database       *mgo.Database
}

func NewMongoDBRepository(collectionName string, DB *mgo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		CollectionName: collectionName,
		Database:       DB,
	}
}
