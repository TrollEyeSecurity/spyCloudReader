package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	mongoPort = "27017"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		os.Getenv("MONGODB_USER"),
		os.Getenv("MONGODB_PASSWD"),
		os.Getenv("MONGODB_HOST"),
		mongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}
	err = MongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return MongoClient, nil
}
