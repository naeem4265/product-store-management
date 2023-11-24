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
	var temp data.Brand
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		fmt.Printf("here %v\n", temp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the MongoDB client
	client, err := CreateMongoDBClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating MongoDB client: %v", err)
		return
	}
	defer client.Disconnect(context.TODO()) // Make sure to close the client when done

	// Access the "brands" collection in the "productStore" database
	collection := client.Database("productStore").Collection("brands")

	// Insert the document into the "brands" collection
	_, err = collection.InsertOne(context.TODO(), temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return a 201 Created status code and the ID of the created document
	w.WriteHeader(http.StatusCreated)
}

func PutBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var temp data.Brand
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

	// Access the "brands" collection in the "productStore" database
	collection := client.Database("productStore").Collection("brands")

	// Create a filter based on the provided ID
	filter := bson.D{{"brand_id", id}}

	// Create an update document using $set operator
	update := bson.D{{"$set", temp}}

	// Update the document in the "brands" collection
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

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
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
