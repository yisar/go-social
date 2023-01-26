package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Novel struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
	Time     string             `json:"time"`
	Sort     string             `json:"sort"`
	Status   string             `json:"status"`
	Size     string             `json:"size"`
	Bio      string             `json:"bio"`
	Tag      string             `json:"tag"`
	Aptitude string             `json:"aptitude"`
	Aid      string             `json:"aid"`
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

func GetNovels(limit, skip *int64) ([]*Novel, error) {
	data := make([]*Novel, 0)
	cursor, err := Mongo.Collection(Novel{}.CollectionName()).
		Find(context.Background(), nil,
			&options.FindOptions{
				Limit: limit,
				Skip:  skip,
				Sort: bson.D{{
					"time", -1,
				}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(Novel)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}