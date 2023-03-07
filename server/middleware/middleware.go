package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
	"github.com/jordandde/supplier-form/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	createDBInstance()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance created")
}

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllSuppliers()
	json.NewEncoder(w).Encode(payload)
}

func CreateSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var supplier models.SupplierList

	err = schema.NewDecoder().Decode(&supplier, r.Form)
	if err != nil {
		log.Fatal(err)
	}

	insertSupplier(supplier)
	json.NewEncoder(w).Encode(supplier)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteSupplier(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func getAllSuppliers() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}}) // sending an empty query to mongodb returns everything in the db
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for curr.Next(context.Background()) {
		var result bson.M
		e := curr.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}
	curr.Close(context.Background())
	return results
}

func deleteSupplier(supplier string) {
	id, _ := primitive.ObjectIDFromHex(supplier)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted 1 supplier")
}

func insertSupplier(supplier models.SupplierList) {

	insertResult, err := collection.InsertOne(context.Background(), supplier)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted a supplier" + fmt.Sprint(insertResult.InsertedID))
}
