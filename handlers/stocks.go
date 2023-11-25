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

func GetStocks(w http.ResponseWriter, r *http.Request) {
	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("productStore").Collection("stocks")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var stocks []bson.M
	for cur.Next(context.TODO()) {
		var stock bson.M
		err := cur.Decode(&stock)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stocks = append(stocks, stock)
	}

	stocksJSON, err := json.MarshalIndent(stocks, "", "  ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(stocksJSON)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("productStore").Collection("stocks")
	filter := bson.D{{"stock_id", id}}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// stock exit or not
	if cur.RemainingBatchLength() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var stocks []bson.M
	for cur.Next(context.TODO()) {
		var stock bson.M
		err := cur.Decode(&stock)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stocks = append(stocks, stock)
	}

	stocksJSON, err := json.MarshalIndent(stocks, "", "  ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(stocksJSON)
}

func PostStocks(w http.ResponseWriter, r *http.Request) {
	var temp data.Stock
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if temp.StockQuantity < 0 {
		log.Printf("Quantity must be positive\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Printf("Error creating MongoDB client: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("productStore").Collection("stocks")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"stock_id": 1},
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

func PutStocks(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	var temp data.Stock
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id != temp.ID {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if temp.StockQuantity < 0 {
		log.Printf("Quantity must be positive\n")
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

	collection := client.Database("productStore").Collection("stocks")
	filter := bson.D{{"stock_id", id}}
	update := bson.D{{"$set", temp}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteStocks(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("stocks")

	filter := bson.D{{"stock_id", id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
