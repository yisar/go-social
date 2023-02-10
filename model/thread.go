package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Thread struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
	Time     string             `json:"time"`
	Sort     string             `json:"sort"`
	Thumb    string             `json:"thumb"`
	Status   string             `json:"status"`
	Size     string             `json:"size"`
	Bio      string             `json:"bio"`
	Tag      string             `json:"tag"`
	Aptitude string             `json:"aptitude"`
	Uid      string             `json:"uid"`
}

func (Thread) CollectionName() string {
	return "thread"
}

func GetThreadByIdentity(identity primitive.ObjectID) (*Thread, error) {
	ub := new(Thread)
	err := Mongo.Collection(Thread{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetThreadCountByName(name string) (int64, error) {
	return Mongo.Collection(Thread{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func InsertThread(thread *Thread) error {
	_, err := Mongo.Collection(Thread{}.CollectionName()).
		InsertOne(context.Background(), bson.D{{"title", thread.Title}, {"content", thread.Content}, {"bio", thread.Bio}, {"status", thread.Status}, {"sort", thread.Sort}, {"size", thread.Size}, {"aptitude", thread.Aptitude}, {"tag", thread.Tag}, {"time", thread.Time}, {"uid", thread.Uid}, {"thumb", thread.Thumb}})
	return err
}

func UpdateThread(thread *Thread, id primitive.ObjectID) error {
	_, err := Mongo.Collection(Thread{}.CollectionName()).
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.D{{"$set",
			bson.D{{"title", thread.Title}, {"content", thread.Content}, {"bio", thread.Bio}, {"status", thread.Status}, {"sort", thread.Sort}, {"size", thread.Size}, {"aptitude", thread.Aptitude}, {"tag", thread.Tag}, {"time", thread.Time}, {"uid", thread.Uid}, {"thumb", thread.Thumb}},
		}})
	return err
}

func GetThreads(limit, skip *int64, sort string) ([]*Thread, error) {
	data := make([]*Thread, 0)
	cursor, err := Mongo.Collection(Thread{}.CollectionName()).
		Find(context.Background(), bson.M{"sort": sort},
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
		mb := new(Thread)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
