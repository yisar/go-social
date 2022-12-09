package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Novel struct {
	Identity  string `json:"_id"`
	Title   string `json:"title"`
	Content  string `json:"content"`
	Time     string `json:"time"`
	Sort    string `json:"sort"`
	Status string `json:"status"`
	Tag string `json:"status"`
	Aid string `json:"aid"`
}

func (Novel) CollectionName() string {
	return "novel"
}

func GetNovelByIdentity(identity primitive.ObjectID) (*Novel, error) {
	ub := new(Novel)
	err := Mongo.Collection(Novel{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetNovelCountByName(name string) (int64, error) {
	return Mongo.Collection(Novel{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func UpdateNovel(novel *Novel) error {
	_, err := Mongo.Collection(Novel{}.CollectionName()).
		InsertOne(context.Background(), novel)
	return err
}