package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Identity  string `bson:"_id"`
	Name   string `bson:"name"`
	Pwd  string `bson:"pwd"`
	Email     string `bson:"email"`
}

func (Author) CollectionName() string {
	return "author"
}

func GetAuthorByAccountPassword(account, password string) (*Author, error) {
	ub := new(Author)
	err := Mongo.Collection(Author{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"name", account}, {"pwd", password}}).
		Decode(ub)
	return ub, err
}

func GetAuthorByIdentity(identity primitive.ObjectID) (*Author, error) {
	ub := new(Author)
	err := Mongo.Collection(Author{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetAuthorCountByEmail(email string) (int64, error) {
	return Mongo.Collection(Author{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}
