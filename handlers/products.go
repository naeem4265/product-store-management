package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/product-store/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		statusNow := product["product_status_id"].(int32)
		active := int32(1)
		if statusNow == active {
			products = append(products, product)
		}
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

func GetProductsById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("productStore").Collection("products")
	filter := bson.D{{"product_id", id}}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// product exit or not
	if cur.RemainingBatchLength() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

	// Check requirement
	collection := client.Database("productStore").Collection("brands")
	filter := bson.D{{"brand_id", temp.BrandID}}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// brand exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Brand id %d not found\n", temp.BrandID)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	collection = client.Database("productStore").Collection("categories")
	filter = bson.D{{"category_id", temp.CategoryID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// category exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Category id %d not found\n", temp.CategoryID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	collection = client.Database("productStore").Collection("suppliers")
	filter = bson.D{{"supplier_id", temp.SupplierID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// supplier exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Supplier id %d not found\n", temp.SupplierID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check product name and product supplier name same or not
	collection = client.Database("productStore").Collection("products")
	filter = bson.D{{"product_supplier_id", temp.SupplierID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		x := product["product_supplier_id"].(int32)
		existingSupplierId := int(x)
		existingProductName := product["product_name"].(string)
		if temp.SupplierID == existingSupplierId && temp.Name == existingProductName {
			log.Printf("Supplier id #%d with name #%s is already exist\n", temp.SupplierID, temp.Name)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
	// everything is ok

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"product_id": 1},
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
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	var temp2 data.Stock
	temp2.ID = temp.ID
	temp2.ProductID = temp.ID
	temp2.StockQuantity = 0
	t := time.Now()
	temp2.UpdatedAt = t.String()
	fmt.Printf("here %v\n", temp2)

	collection = client.Database("productStore").Collection("stocks")

	indexModel = mongo.IndexModel{
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

	_, err = collection.InsertOne(context.TODO(), temp2)
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

func PutProducts(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	var temp data.Product
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id != temp.ID {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// check brand, supplier, category id
	collection := client.Database("productStore").Collection("brands")
	filter := bson.D{{"brand_id", temp.BrandID}}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// brand exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Brand id %d not found\n", temp.BrandID)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	collection = client.Database("productStore").Collection("categories")
	filter = bson.D{{"category_id", temp.CategoryID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// category exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Category id %d not found\n", temp.CategoryID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	collection = client.Database("productStore").Collection("suppliers")
	filter = bson.D{{"supplier_id", temp.SupplierID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())
	// supplier exit or not
	if cur.RemainingBatchLength() == 0 {
		log.Printf("Supplier id %d not found\n", temp.SupplierID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check product name and product supplier name same or not
	collection = client.Database("productStore").Collection("products")
	filter = bson.D{{"product_supplier_id", temp.SupplierID}}
	cur, err = collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		x := product["product_supplier_id"].(int32)
		existingSupplierId := int(x)
		existingProductName := product["product_name"].(string)
		if temp.SupplierID == existingSupplierId && temp.Name == existingProductName {
			log.Printf("Supplier id #%d with name #%s is already exist\n", temp.SupplierID, temp.Name)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
	// everything ok

	collection = client.Database("productStore").Collection("products")
	filter = bson.D{{"product_id", id}}
	update := bson.D{{"$set", temp}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	client, err := CreateMongoDBClient()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("productStore").Collection("products")

	filter := bson.D{{"product_id", id}}
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
