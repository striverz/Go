package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/striverz/100xGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://striverz:Kmk%400011@100xdevs.06113os.mongodb.net/"
const databaseName = "Netflix"
const collectionName = "Watchlist"

var collection *mongo.Collection

func init() {
	
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}


//insert One Record
func insertOneMovie(movie models.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted One movie in db with it ",inserted.InsertedID)
}


//update One Record
func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)

	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	

	result,err:=collection.UpdateOne(context.Background(),filter,update)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("modified count : ",result.ModifiedCount)
}

//delete One Record

func deleteOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)

	filter:=bson.M{"_id":id}

	result,err:=collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("deletion result is: ",result.DeletedCount)

}

func deleteAllMovies(){
	result,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("deletion all result is ",result.DeletedCount)
}

func getAllMovies() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

func GetAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	allMovies:=getAllMovies()

	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie models.Netflix

	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(map[string]interface{}{
    "message": "Movie is Created",
    "movie":   movie,
})

}


func DeleMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	deleteOneMovie(chi.URLParam(r, "id"))

	json.NewEncoder(w).Encode(chi.URLParam(r,"id"))
}