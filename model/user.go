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
	Level    int                `json:"level"`
}

func (User) CollectionName() string {
	return "user"
}

func GetUserByAccountPassword(name, pwd string) (*User, error) {
	fmt.Println(name, pwd)
	ub := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"name", name}, {"pwd", pwd}}).
		Decode(ub)
	return ub, err
}

func GetUserByIdentity(identity primitive.ObjectID) (*User, error) {
	ub := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}

func GetUserCountByEmail(email string) (int64, error) {
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

func GetUserCountByName(name string) (int64, error) {
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

func InsertUser(user *User) error {
	_, err := Mongo.Collection(User{}.CollectionName()).
		InsertOne(context.Background(), bson.D{{"name", user.Name}, {"pwd", user.Pwd}, {"email", user.Email}, {"level", 0}})
	return err
}
func UpdateUser(user *User, id primitive.ObjectID) error {
	_, err := Mongo.Collection(User{}.CollectionName()).
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.D{{"$set", bson.D{{"name", user.Name}, {"pwd", user.Pwd}, {"email", user.Email}, {"level", 0}}}})
	return err
}
