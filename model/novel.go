package model

import (
	"context"
	"fmt"
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

func GetNovels() []*Novel {
	cur, err := Mongo.Collection(Novel{}.CollectionName()).Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Printf("查询失败，err=%v \n", err)
		return nil
	}

	var list []*Novel

	// list := make(map[string]*Novel)

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {

		novel := Novel{}

		err := cur.Decode(novel)
		if err != nil {
			fmt.Printf("解析失败 err=%v \n", err)
			continue
		}

		list = append(list, &novel)

	}

	return list
}
