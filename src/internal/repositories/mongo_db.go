package repositories

import (
	"Faceit/src/internal/model"
	"context"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
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

func (repo *MongoDBRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := repo.Database.C(repo.CollectionName).Insert(user)
	return user, err
}

func (repo *MongoDBRepository) UpdateUser(ctx context.Context, userId string, user *model.User) (*model.User, error) {
	existingUser, err := repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": bson.ObjectIdHex(userId)}
	existingUser.Merge(*user)
	update := bson.M{
		"$set": existingUser,
	}

	err = repo.Database.C(repo.CollectionName).Update(filter, update)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (repo *MongoDBRepository) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user *model.User
	filter := bson.M{"id": bson.ObjectIdHex(userId)}
	err := repo.Database.
		C(repo.CollectionName).
		Find(filter).
		One(&user)

	return user, err
}

func (repo *MongoDBRepository) DeleteUser(ctx context.Context, userId string) error {
	filter := bson.M{"id": bson.ObjectIdHex(userId)}
	err := repo.Database.C(repo.CollectionName).Remove(filter)

	return err
}

func (repo *MongoDBRepository) GetUsers(ctx context.Context, key, value string) ([]model.User, error) {
	var usersArray []model.User
	if key == "first_name" || key == "last_name" || key == "nickname" || key == "password" || key == "email" || key == "country" {
		filter := bson.M{key: value}
		err := repo.Database.C(repo.CollectionName).Find(filter).All(&usersArray)
		if err != nil {
			return nil, err
		}
	}

	return usersArray, nil
}
