package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chapter struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Oid      int                `json:"oid"`
	Nid      string             `json:"nid"`
	Status   string             `json:"status"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
	Time     string             `json:"time"`
}

type Chapter2 struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Oid      int                `json:"oid"`
	Nid      string             `json:"nid"`
	Status   string             `json:"status"`
	Title    string             `json:"title"`
	Time     string             `json:"time"`
}

func (Chapter) CollectionName() string {
	return "chapter"
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

func InsertChapter(chapter *Chapter) error {
	_, err := Mongo.Collection(Chapter{}.CollectionName()).
		InsertOne(context.Background(), bson.D{{"oid", chapter.Oid}, {"title", chapter.Title}, {"content", chapter.Content}, {"status", chapter.Status}, {"time", chapter.Time}, {"nid", chapter.Nid}})
	return err
}

func UpdateChapter(chapter *Chapter, id primitive.ObjectID) error {
	_, err := Mongo.Collection(Chapter{}.CollectionName()).
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.D{{"oid", chapter.Oid}, {"title", chapter.Title}, {"content", chapter.Content}, {"status", chapter.Status}, {"time", chapter.Time}, {"nid", chapter.Nid}})
	return err
}

func GetChapters(limit, skip *int64, nid string) ([]*Chapter2, error) {
	data := make([]*Chapter2, 0)
	cursor, err := Mongo.Collection(Chapter{}.CollectionName()).
		Find(context.Background(), bson.M{"nid": nid},
			&options.FindOptions{
				Limit: limit,
				Skip:  skip,
				Sort: bson.D{{
					"oid", 1,
				}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(Chapter2)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
