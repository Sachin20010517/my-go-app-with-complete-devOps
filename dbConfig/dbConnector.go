package dbConfig

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() {

	//printing a message to the console
	fmt.Println("Connecting to mongo cluster")

	//Contexts are essential in Go for managing timeouts, cancellations, and passing request-scoped data	
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}

	//Ping Database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to mongo cluster")
	DATABASE = client.Database(DATABASE_NAME)

}