package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Sign     string             `json:"sign"`
	Location []int              `json:"location"`
	Level    int                `json:"level"`
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
			{"sign", user.Sign},
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
			{"sign", user.Sign},
			{"location", user.Location},
			{"level", user.Level},
		})
	return err
}
