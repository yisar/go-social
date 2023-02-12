package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Summary  string             `json:"summary"`
	Tid      string             `json:"tid"`
	Status   string             `json:"status"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
	Time     string             `json:"time"`
	Length   int                `json:"length"`
	Uname    string             `json:"uname"`
}

type Post2 struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Summary  string             `json:"summary"`
	Tid      string             `json:"tid"`
	Status   string             `json:"status"`
	Title    string             `json:"title"`
	Time     string             `json:"time"`
	Length   int                `json:"length"`
	Uname    string             `json:"uname"`
}

func (Post) CollectionName() string {
	return "post"
}

func GetPostByIdentity(identity primitive.ObjectID) (*Post, error) {
	ub := new(Post)
	err := Mongo.Collection(Post{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetPostCountByName(name string) (int64, error) {
	return Mongo.Collection(Post{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func InsertPost(post *Post) error {
	_, err := Mongo.Collection(Post{}.CollectionName()).
		InsertOne(context.Background(), bson.D{{"summary", post.Summary}, {"uname", post.Uname}, {"title", post.Title}, {"content", post.Content}, {"status", post.Status}, {"time", post.Time}, {"tid", post.Tid}, {"length", post.Length}})
	return err
}

func UpdatePost(post *Post, id primitive.ObjectID) error {
	_, err := Mongo.Collection(Post{}.CollectionName()).
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.D{{"summary", post.Summary}, {"uname", post.Uname}, {"title", post.Title}, {"content", post.Content}, {"status", post.Status}, {"time", post.Time}, {"tid", post.Tid}, {"length", post.Length}})
	return err
}

func GetPosts(limit, skip *int64, tid string) ([]*Post2, error) {
	data := make([]*Post2, 0)
	cursor, err := Mongo.Collection(Post{}.CollectionName()).
		Find(context.Background(), bson.M{"tid": tid},
			&options.FindOptions{
				Limit: limit,
				Skip:  skip,
				Sort: bson.D{{
					"time", 1,
				}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(Post2)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
