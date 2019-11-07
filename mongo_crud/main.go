package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database       = "test_db"
	collection     = "post"
	replicaSetName = "mongo-rs"
)

var Client *mongo.Client

type Post struct {
	Tile    string `json:"tile" bson:"tile"`
	Content string `json:"content" bson:"content"`
}

func connect() {
	// uri := "mongodb://mongoadmin:secret@localhost:27017"
	// change stream work only with replica set
	uri := "mongodb+srv://mongoadmin:secret1234@cluster0-xxyrd.gcp.mongodb.net"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}

func AddAndSleep() {
	for i := 0; i < 10; i++ {
		AddOne()
		log.Println("sleep 300ms")
		time.Sleep(time.Millisecond * 300)
	}
}

func AddOne() {
	res, err := Client.Database(database).Collection(collection).InsertOne(context.TODO(), Post{"Vietnam thang Indo", "Viet Nam thang dam Indo 3-1"})
	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	log.Println(id)
}

func main() {

	connect()
	// AddOne()
	AddAndSleep()
}
