package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongo-golang-hitesh/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString ="mongodb+srv://arpit:arpit1234@cluster0.emlym8n.mongodb.net/?retryWrites=true&w=majority"
const dbName ="netflix"
const colName="watchList"

//Most Important
var collection *mongo.Collection


func init(){

	// Client Option
	var clientOption = options.Client().ApplyURI(connectionString)

	// Connection
	client,err :=mongo.Connect(context.TODO(),clientOption)

	if err!=nil{

		panic(err)
	}

	fmt.Println("Connected to Mongodb")

	collection=client.Database(dbName).Collection(colName)

	fmt.Println("Collection Instance is ready")

}


//MongoDB Helpers -- should be in separate file but for now we are making them here

// insert Single movie
func insertOneMovie(movie model.Netflix){

	inserted,err:=collection.InsertOne(context.TODO(),movie)

	if(err!=nil){

		fmt.Println("Error While Inserting a movie")
		fmt.Println(err)
		return
	}

	fmt.Println("Movie inserted with id : ",inserted.InsertedID)
}

//update one record

func updateOneRecord(movieId string){

	id,_:=primitive.ObjectIDFromHex(movieId)

	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}

	result,err:=collection.UpdateOne(context.Background(),filter,update)

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("No. of Records Updated are : ",result.ModifiedCount)
}


//delete 1 Record

func deleteOneMovie(movieId string) int{

	id,_:=primitive.ObjectIDFromHex(movieId)

	filter:=bson.M{"_id":id}

	result,err:=collection.DeleteOne(context.Background(),filter)
	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("No. of Records Deleted are : ",result.DeletedCount)
	return int(result.DeletedCount)
}

//Delete all Record

func deleteAllMovie() int{

	 result,err:=collection.DeleteMany(context.Background(),bson.M{})
	 if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("No. of Records Deleted are : ",result.DeletedCount)
	return int(result.DeletedCount)
}

// Get all Movies

func getAllMovies() []model.Netflix{

	var movies []model.Netflix
	cursor , err:= collection.Find(context.Background(),bson.M{})
	if(err!=nil){
		log.Fatal(err)
	}

	for cursor.Next(context.Background()){

		var movie model.Netflix

		err:=cursor.Decode(&movie)
		if(err!=nil){
			log.Fatal(err)
		}
		movies=append(movies,movie)
	}

	defer cursor.Close(context.Background())
	return movies
	// cur, err := collection.Find(context.Background(), bson.D{{}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var movies []primitive.M

	// for cur.Next(context.Background()) {
	// 	var movie bson.M
	// 	err := cur.Decode(&movie)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	movies = append(movies, movie)
	// }

	// defer cur.Close(context.Background())
	// return movies
}



// Actual Controller function(ones which we will use in router file)

func GetAllMovies(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(getAllMovies())
}

func CreateMovie(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	fmt.Println(r.Body)

	var movie model.Netflix

	_=json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params:=mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params:=mux.Vars(r)
	deleteCount:=deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(deleteCount)

}

func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	deleteCount:=deleteAllMovie()
	json.NewEncoder(w).Encode(deleteCount)


}

