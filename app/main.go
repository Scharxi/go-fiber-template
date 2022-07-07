package main

import (
	"context"
	"fiber-app-template/app/database"
	"fiber-app-template/app/server"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mongoClient := database.Connect()
	db := mongoClient.Database("mydb")

	defer func() {
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("The database connection has been successfully closed.")
	}()

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err := database.Mongoose.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln("The database do not have an open connection.")
	} else {
		log.Println("The database has been started successfully.")
	}

	// creating and starting the server.
	port := os.Getenv("PORT")
	server.Connect(port, *db)
}
