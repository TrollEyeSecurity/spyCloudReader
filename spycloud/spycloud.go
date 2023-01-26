package spycloud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"spyCloudReader/mongodb"
)

func UpdateOrCreateBreachRecordFromString(breachRecordObjJson *BreachRecordMongoDb) {
	collection := fmt.Sprintf("SpyCloudDataLake")
	MongoClient, MongoClientError := mongodb.GetMongoClient()
	defer MongoClient.Disconnect(context.TODO())
	if MongoClientError != nil {
		err := fmt.Errorf("update-or-create breach record error %v", MongoClientError)
		fmt.Println(err)
		log.Println(err)
		return
	}
	breachRecordCollection := MongoClient.Database(os.Getenv("MONGODB_DB")).Collection(collection)
	var breachRecordObjResult bson.M
	breachRecordObjJson.Id = breachRecordObjJson.DocumentId
	findOneErr := breachRecordCollection.FindOne(context.TODO(), bson.D{{"_id", breachRecordObjJson.Id}}).Decode(&breachRecordObjResult)
	if findOneErr != nil {
		if findOneErr == mongo.ErrNoDocuments {
			_, insertOneErr := breachRecordCollection.InsertOne(context.TODO(), breachRecordObjJson)
			if insertOneErr != nil {
				err := fmt.Errorf("insert breach record error %v", findOneErr)
				fmt.Println(err)
				log.Println(err)
				MongoClient.Disconnect(context.TODO())
				return
			}
		} else {
			err := fmt.Errorf("find breach record error %v", findOneErr)
			fmt.Println(err)
			log.Println(err)
			MongoClient.Disconnect(context.TODO())
			return
		}
	} else {
		_, updateOneErr := breachRecordCollection.UpdateOne(context.TODO(), bson.D{{"_id", breachRecordObjJson.Id}}, bson.D{{"$set", *breachRecordObjJson}})
		if updateOneErr != nil {
			err := fmt.Errorf("update host error %v", updateOneErr)
			fmt.Println(err)
			log.Println(err)
			MongoClient.Disconnect(context.TODO())
			return
		}
	}
}
