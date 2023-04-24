package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Identity primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Pwd      string             `json:"pwd"`
	Email    string             `json:"email"`
	Age      int                `json:"age"`
	Sex      int                `json:"sex"`
	Height   int                `json:"height"`
	Weight   int                `json:"weight"`
	Bio      string             `json:"bio"`
	Location []float64          `json:"location"`
	Level    int                `json:"level"`
	Distance float64 `json:"distance"`
}

func GetUserByAccountPassword(name, pwd string) (*User, error) {
	fmt.Println(name, pwd)
	ub := new(User)
	err := Mongo.Collection("user").
		FindOne(context.Background(), bson.D{{"name", name}, {"pwd", pwd}}).
		Decode(ub)
	return ub, err
}

func GetUserByIdentity(identity primitive.ObjectID) (*User, error) {
	ub := new(User)
	err := Mongo.Collection("user").
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetUserCountByEmail(email string) (int64, error) {
	return Mongo.Collection("user").
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

func GetUserCountByName(name string) (int64, error) {
	return Mongo.Collection("user").
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func InsertUser(user *User) error {
	_, err := Mongo.Collection("user").
		InsertOne(context.Background(), bson.D{
			{"name", user.Name},
			{"pwd", user.Pwd},
			{"email", user.Email},
			{"age", user.Age},
			{"sex", user.Sex},
			{"height", user.Height},
			{"weight", user.Weight},
			{"bio", user.Bio},
			{"location", user.Location},
			{"level", 1},
		})
	return err
}
func UpdateUser(user *User, id primitive.ObjectID) error {
	_, err := Mongo.Collection("user").
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.D{
			{"name", user.Name},
			{"pwd", user.Pwd},
			{"email", user.Email},
			{"age", user.Age},
			{"sex", user.Sex},
			{"height", user.Height},
			{"weight", user.Weight},
			{"bio", user.Bio},
			{"location", user.Location},
			{"level", user.Level},
		})
	return err
}

func GetUsers(location []float64) ([]*User, error) {
	data := make([]*User, 0)
	stages := mongo.Pipeline{}
	getNearbyStage := bson.D{{"$geoNear", bson.M{
		"near":               location,
		"distanceMultiplier": 6378137,
		"maxDistance":        1,
		"spherical":          true,
		"distanceField":      "distance",
	}}}
	stages = append(stages, getNearbyStage)
	cursor, err := Mongo.Collection("user").Aggregate(context.Background(), stages)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(User)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
