package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Identity  primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string `json:"name"`
	Pwd  string `json:"pwd"`
	Email     string `json:"email"`
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

func GetAuthorCountByName(name string) (int64, error) {
	return Mongo.Collection(Author{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func UpdateAuthor(author *Author) error {
	_, err := Mongo.Collection(Author{}.CollectionName()).
		InsertOne(context.Background(), bson.D{{"name", author.Name}, {"pwd", author.Pwd},{"email", author.Email}})
	return err
}