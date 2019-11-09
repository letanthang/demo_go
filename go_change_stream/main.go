package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var Client *mongo.Client

func main() {

	connect()

	collection := Client.Database("test_db").Collection("post")

	ctx := context.Background()
	stream, err := collection.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		log.Println(err)
		return
	}
	defer stream.Close(ctx)

	for stream.Next(ctx) {
		var elem bson.M

		if err := stream.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Println(elem)
	}

}

func connect() {
	// uri := "mongodb://mongoadmin:mongosecret@localhost:27017"
	uri := "mongodb+srv://mongoadmin:secret1234@cluster0-xxyrd.gcp.mongodb.net"

	fmt.Println("Connect MongoDb with uri", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Client = client

}
