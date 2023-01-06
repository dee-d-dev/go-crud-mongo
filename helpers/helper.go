package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dee-d-dev/go-mongodb-crud/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

var connectionString string
var dbName string


var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	err := godotenv.Load() // 👈 load .env file
    if err != nil {
    	log.Fatal(err)
    }
    
	connectionString = os.Getenv("connectionString")
	dbName = os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")
    
    // fmt.Println(connectionString)
    

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection is ready")
}

func CreateTodo(todo models.Todo) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo added")

	return inserted
}

func UpdateTodo(todoId string) *mongo.UpdateResult {
	id, _ := primitive.ObjectIDFromHex(todoId)

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"completed": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated", id)
	return result
}

func GetAllTodos() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var todos []primitive.M

	for cursor.Next(context.Background()) {
		var todo bson.M
		err := cursor.Decode(&todo)

		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	defer cursor.Close(context.Background())
	return todos
}

func DeleteTodo(todoId string) *mongo.DeleteResult {
	id, _ := primitive.ObjectIDFromHex(todoId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", id)
	return result
}

func DeleteAll() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount

}
