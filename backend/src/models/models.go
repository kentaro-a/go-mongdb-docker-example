package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	HOST string = "mongodb://db:27017"
	USER string = "root"
	PASS string = "mongo"
)

func GetConnection() *mongo.Client {
	cred := &options.Credential{
		Username: USER,
		Password: PASS,
	}
	opt := options.Client()
	opt.Auth = cred
	opt = opt.ApplyURI(HOST)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

type User struct {
	ID   int
	Name string
	Pass string
}

func GetUserByNamePass(name string, pass string) *User {
	conn := GetConnection()
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := conn.Database("test").Collection("users")

	filter := bson.D{{"name", name}, {"pass", pass}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := &User{}
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil
	}
	return user
}
