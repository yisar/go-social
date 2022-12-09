package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chapter struct {
	Identity  string `json:"_id"`
	Title   string `json:"title"`
	Content  string `json:"content"`
	Time     string `json:"time"`
}

func (Chapter) CollectionName() string {
	return "Chapter"
}

func GetChapterByIdentity(identity primitive.ObjectID) (*Chapter, error) {
	ub := new(Chapter)
	err := Mongo.Collection(Chapter{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetChapterCountByName(name string) (int64, error) {
	return Mongo.Collection(Chapter{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func UpdateChapter(chapter *Chapter) error {
	_, err := Mongo.Collection(Chapter{}.CollectionName()).
		InsertOne(context.Background(), chapter)
	return err
}