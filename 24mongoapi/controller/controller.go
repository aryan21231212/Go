package controller

import (
	"context"
	"fmt"
	"github/aryan/mongodb/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://aryannoob2121:hvD1ncMp8oQLMu8z@cluster0.dkwjh7f.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// insert one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.insertone(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One record successfully Inserted of ID: ", inserted.insertedID)
}
