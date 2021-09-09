package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/d-kuznetsov/calendar-backend/models"
)

type mongoUser struct {
	Id       primitive.ObjectID `bson:"id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func toModelUser(user mongoUser) models.User {
	return models.User{
		Id:       user.Id.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func CreateMongoClient(uri string) *mongo.Client {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

func CreateMongoRepo(client *mongo.Client, dbName string) IRepository {
	return &MongoRepository{
		client: client,
		dbName: dbName,
	}
}

type MongoRepository struct {
	client *mongo.Client
	dbName string
}

func (repo *MongoRepository) CreateUser(name, email, hashedPassword string) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := mongoUser{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (repo *MongoRepository) GetUserByEmail(email string) (models.User, error) {
	var user mongoUser
	collection := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, ErrNoUsersFound
	}
	return toModelUser(user), err
}
