package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github/aryan/mongodb/model"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aryannoob2121:hvD1ncMp8oQLMu8z@cluster0.dkwjh7f.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// 🔹 Initialize MongoDB connection
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

// 🔹 Insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie) // ✅ corrected method name
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One record successfully inserted with ID:", inserted.InsertedID) // ✅ fixed field name case
}

// 🔹 Update one movie
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

// 🔹 Delete one movie
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter) // ✅ removed nil param
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted count:", deleteResult.DeletedCount) // ✅ corrected print variable
}

// 🔹 Delete all movies
func deleteManyMovie() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteResult.DeletedCount)
}

// 🔹 Get all movies
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background()) // ✅ always defer after successful cursor creation

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// ===============================================================
// 🔹 Controller functions (these are handlers for HTTP routes)
// ===============================================================

// ✅ Get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // ✅ fixed content type
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// ✅ Create movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // ✅ corrected header name

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// ✅ Mark movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT") // ✅ corrected header name

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// ✅ Delete one movie
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // ✅ corrected header name

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Deleted the movie")
}

// ✅ Delete all movies
func DeleteManyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // ✅ corrected header name

	deleteManyMovie()
	json.NewEncoder(w).Encode("All movies are deleted!")
}
