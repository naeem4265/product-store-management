package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/product-store/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBrands(w http.ResponseWriter, r *http.Request) {
	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO()) // Make sure to close the client when done
	// Access the "brands" collection in the "productStore" database

	collection := client.Database("productStore").Collection("brands")
	// Get the documents from the "brands" collection
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO()) // Close the cursor when done

	// Create a slice to hold the retrieved documents
	var brands []bson.M
	for cur.Next(context.TODO()) {
		var brand bson.M
		err := cur.Decode(&brand)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		brands = append(brands, brand)
	}

	// Convert the brands slice to JSON
	brandsJSON, err := json.MarshalIndent(brands, "", "  ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(brandsJSON)
}

func PostBrand(w http.ResponseWriter, r *http.Request) {
	// get data from json
	var temp data.Brand
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Printf("Error creating MongoDB client: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	// Access the "brands" collection in the "productStore" database
	collection := client.Database("productStore").Collection("brands")

	// Create a unique compound index on the "brand_id" and "brand_name" fields
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"brand_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the document into the "brands" collection
	_, err = collection.InsertOne(context.TODO(), temp)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Printf("Duplicate Brand Id. Error inserting document: %v\n", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PutBrand(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	var temp data.Brand
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id != temp.BrandId {
		w.WriteHeader(http.StatusConflict)
		return
	}
	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("brands")

	// Create a filter based on the provided ID
	filter := bson.D{{"brand_id", id}}
	// Create an update document using $set operator
	update := bson.D{{"$set", temp}}

	// Update the document in the "brands" collection
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if a document was updated
	if result.ModifiedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO()) // Make sure to close the client when done

	// Access the "brands" collection in the "productStore" database
	collection := client.Database("productStore").Collection("brands")

	// Create a filter based on the provided ID
	filter := bson.D{{"brand_id", id}}
	// Delete the document from the "brands" collection
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if a document was deleted
	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
