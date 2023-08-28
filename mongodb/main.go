package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	urlMongo := os.Getenv("MONGO_URI")
	// Estabeleça uma conexão com o MongoDB
	clientOptions := options.Client().ApplyURI(urlMongo)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Selecione um banco de dados
	db := client.Database("test")

	// Liste as coleções
	colls, err := db.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Imprime os nomes das coleções
	for _, coll := range colls {
		fmt.Println(coll)
	}
}
