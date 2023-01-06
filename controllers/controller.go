package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dee-d-dev/go-mongodb-crud/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson"
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

let connectionString
let dbName 
let todos

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	err := gotdotenv.Load() // 👈 load .env file
    if err != nil {
    	log.Fatal(err)
    }
    
	connectionString := os.Getenv("connectionString")
	dbName := os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")
    
    fmt.Println(connectionString)
    

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection is ready")
}

func createTodo(todo models.Todo) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(context.Background(), todo)

	helper.HandleError(err)
	fmt.Println("Todo added")

	return inserted
}

func updateTodo(todoId string) *mongo.UpdateResult {
	id, _ = primitive.ObjectIDFromHex(todoId)

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"completed": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	helper.HandleError(err)

	fmt.Println("updated", id)
	return result
}

func getAllTodos() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	helper.HandleError(err)

	var todos []primitive.M

	for cursor.Next(context.Background()) {
		var todo bson.M
		err := cursor.Decode(&todo)

		helper.HandleError(err)
		todos = append(todos, todo)
	}

	defer cursor.close(context.Background())
	return todos
}

func deleteTodo(todoId string) *mongo.DeleteResukf {
	id, _ = primitive.ObjectIDFromHex(todoId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	helper.HandleError(err)
	fmt.Println("Deleted", id)
	return result
}

func deleteAll() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	helper.HandleError(err)

	return result.DeletedCount

}

func CreateOne(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	result := createTodo(todo)

	json.NewEncoder(w).Encode(result)
}


func UpdateTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	result := updateTodo(params["id"])
	json.NewEncoder(w).Encode(result)

}

func Delete(w http.ReponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := deleteTodo(params["id"])
	json.NewEncoder(w).Encode(result)
}

func Todos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	result := getAllTodos()
	json.NewEncoder(w).Encode(result)
}

func DeleteEverything(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

    result := deleteAll()
    json.NewEncoder(w).Encode(result)
}