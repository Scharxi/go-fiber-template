package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongoose *mongo.Client
var err error

func Connect() *mongo.Client {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "admin",
		Username:      "root",
		Password:      "secret",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:50030/").SetAuth(credential)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Mongoose, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalln(err)
	}

	return Mongoose
}
