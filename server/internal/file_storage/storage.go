package file_storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

const username = "admin"
const password = "password"
const database = "file"

var uri = fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s?authSource=admin", username, password, database)

var client, _ = mongo.Connect(options.Client().ApplyURI(uri))
var ctx = context.TODO()

type fileDal struct {
	Filename string `bson:"filename"`
	Data     []byte `bson:"data"`
}

func SaveFile(filename string, file []byte) (string, error) {
	collection := client.Database("file").Collection("audio")

	result, err := collection.InsertOne(ctx, bson.M{
		"filename": filename,
		"data":     file,
	})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	insertedID := result.InsertedID.(bson.ObjectID).Hex()

	return insertedID, nil
}

func DeleteFile(id string) error {
	objectId, objectIdErr := bson.ObjectIDFromHex(id)

	if objectIdErr != nil {
		return fmt.Errorf("invalid id")
	}

	collection := client.Database("file").Collection("audio")
	_, deleteErr := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func GetFile(id string) ([]byte, string, error) {
	objectId, _ := bson.ObjectIDFromHex(id)
	collection := client.Database("file").Collection("audio")

	var result fileDal
	err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result)

	fmt.Println(result.Filename)

	if err != nil {
		return nil, "", err
	}

	return result.Data, result.Filename, nil
}
