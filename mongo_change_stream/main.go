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
	ID            IDELem          `json:"id" bson:"_id"`
	OperationType string          `json:"operationType" bson:"operationType"`
	FullDocument  Post            `json:"fullDocument" bson:"fullDocument"`
	NS            NSELem          `json:"ns" bson:"ns"`
	DocumentKey   DocumentKeyElem `json:"documentKey" bson:"documentKey"`
}

func main() {
	uri := "mongodb://mongoadmin:secret@localhost:27017/test_db"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	col := client.Database(database).Collection(collection)

	var pipeline interface{}
	ctx := context.Background()
	stream, err := col.Watch(ctx, pipeline)
	if err != nil {
		return
	}
	defer stream.Close(ctx)

	for stream.Next(ctx) {
		elem := CSElem{}

		if err := stream.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Println(elem)
	}

	if err := stream.Err(); err != nil {
		log.Fatal(err)
	}
}
