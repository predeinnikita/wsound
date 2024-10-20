package file_storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"io/ioutil"
	"log"
)

var client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:32770"))
var ctx = context.TODO()

func SaveFile() {
	data, err := ioutil.ReadFile("./test.mp3")
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("file").Collection("audio")

	mp3Document := bson.M{
		"filename": "file.mp3",
		"data":     data,
	}

	_, err = collection.InsertOne(ctx, mp3Document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MP3 файл успешно загружен!")
}
