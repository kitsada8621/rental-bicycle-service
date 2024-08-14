package configs

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDb() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("DB_URI")
	fmt.Println("DB URI: ", uri)

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return Client, nil

}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	database := os.Getenv("DB_NAME")
	fmt.Println("DB_NAME", database)
	return client.Database(database).Collection(collectionName)
}
