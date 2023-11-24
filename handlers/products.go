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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("products")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	var products []bson.M
	for cur.Next(context.TODO()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	productsJSON, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(productsJSON)
}

func PostProducts(w http.ResponseWriter, r *http.Request) {
	var temp data.Product
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
	collection := client.Database("productStore").Collection("products")
	_, err = collection.InsertOne(context.TODO(), temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PutProducts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var temp data.Product
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("products")
	filter := bson.D{{"product_id", id}}
	update := bson.D{{"$set", temp}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("products")
	filter := bson.D{{"product_id", id}}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
