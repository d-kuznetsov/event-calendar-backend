package mongodb

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
	"github.com/d-kuznetsov/calendar-backend/repository"
)

func CreateClient(uri string) *mongo.Client {
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

type dbUser struct {
	Id       primitive.ObjectID `bson:"id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func toModelUser(user dbUser) models.User {
	return models.User{
		Id:       user.Id.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (repo *mongoRepo) CreateUser(name, email, hashedPassword string) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := dbUser{
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

func (repo *mongoRepo) GetUserByEmail(email string) (models.User, error) {
	var user dbUser
	collection := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, repository.ErrNoUsersFound
	}
	return toModelUser(user), err
}

type dbEvent struct {
	Id        primitive.ObjectID `bson:"id,omitempty"`
	Date      string             `json:"date"`
	StartTime string             `json:"startTime"`
	EndTime   string             `json:"endTime"`
	Content   string             `json:"content"`
	UserId    primitive.ObjectID `json:"userId"`
}

func (repo *mongoRepo) CreateEvent(params repository.EventOpts) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userId, err := primitive.ObjectIDFromHex(params.UserId)
	if err != nil {
		return "", err
	}
	event := dbEvent{
		Date:      params.Date,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Content:   params.Content,
		UserId:    userId,
	}
	res, err := coll.InsertOne(ctx, event)
	if err != nil {
		return "", err
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}
