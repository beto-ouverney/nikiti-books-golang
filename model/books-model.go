package model

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection

func init() {

	var cred options.Credential

	cred.Username = config.MONGO_USER
	cred.Password = config.MONGO_PASSWORD

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MONGO_CONNECT).SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	collection = client.Database("nikiti_books_db").Collection("books")
	fmt.Println("Collection is ready")

}
