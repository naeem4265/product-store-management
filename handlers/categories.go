package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/product-store/data"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("categories")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	var categories []bson.M
	for cur.Next(context.TODO()) {
		var category bson.M
		err := cur.Decode(&category)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	categoriesJSON, err := json.MarshalIndent(categories, "", "  ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(categoriesJSON)
}

func PostCategories(w http.ResponseWriter, r *http.Request) {
	var temp data.Category
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		fmt.Printf("here %v\n", temp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	client, err := CreateMongoDBClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating MongoDB client: %v", err)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("categories")
	_, err = collection.InsertOne(context.TODO(), temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PutCategories(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var temp data.Category
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO()) // Make sure to close the client when done

	// Access the "categories" collection in the "productStore" database
	collection := client.Database("productStore").Collection("categories")

	// Create a filter based on the provided ID
	filter := bson.D{{"category_id", id}}

	// Create an update document using $set operator
	update := bson.D{{"$set", temp}}

	// Update the document in the "categories" collection
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO()) // Make sure to close the client when done

	// Access the "categories" collection in the "productStore" database
	collection := client.Database("productStore").Collection("categories")

	// Create a filter based on the provided ID
	filter := bson.D{{"category_id", id}}

	// Delete the document from the "categories" collection
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
