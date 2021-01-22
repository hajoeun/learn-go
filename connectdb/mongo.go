package connectdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo connect with mongo db
func Mongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://joeun:<password>@cluster0.kypvf.mongodb.net",
	))
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("sample_mflix").Collection("movies")

	opts := options.Find().SetSort(bson.D{{"year", 1}})
	cursor, err := coll.Find(context.TODO(), bson.D{{"year", 1909}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
