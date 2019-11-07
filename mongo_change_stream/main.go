package main

import (
	"context"
	"log"

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
type IDELem struct {
	Data string `json:"data" bson:"_data"`
}
type NSELem struct {
	DB   string `json:"db" bson:"db"`
	Coll string `json:"coll" bson:"coll"`
}
type DocumentKeyElem struct {
	ID string `json:"id" bson:"_id"`
}
type CSElem struct {
	ID            IDELem      `json:"id" bson:"_id"`
	OperationType string      `json:"operationType" bson:"operationType"`
	FullDocument  Post        `json:"fullDocument" bson:"fullDocument"`
	NS            NSELem      `json:"ns" bson:"ns"`
	DocumentKey   interface{} `json:"documentKey" bson:"documentKey"`
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

func main() {

	connect()
	col := Client.Database(database).Collection(collection)

	ctx := context.Background()
	stream, err := col.Watch(ctx, mongo.Pipeline{}) // bson.A{}
	if err != nil {
		log.Println(err)
		return
	}
	defer stream.Close(ctx)

	for stream.Next(ctx) {
		elem := CSElem{}

		if err := stream.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Printf("event %+v \n", elem)
	}

	if err := stream.Err(); err != nil {
		log.Fatal(err)
	}
}
