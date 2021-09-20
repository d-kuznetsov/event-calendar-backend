package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/d-kuznetsov/calendar-backend/repository"
)

var timeout = 10 * time.Second

func CreateClient(uri string) *mongo.Client {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB is open.")
	return client
}

type mongoRepo struct {
	client *mongo.Client
	dbName string
}

func CreateRepository(client *mongo.Client, dbName string) repository.IRepository {
	return &mongoRepo{
		client: client,
		dbName: dbName,
	}
}
