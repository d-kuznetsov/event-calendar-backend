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

	"github.com/d-kuznetsov/calendar-backend/entities"
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

func toModelUser(user dbUser) entities.User {
	return entities.User{
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

func (repo *mongoRepo) GetUserByEmail(email string) (entities.User, error) {
	var user dbUser
	collection := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return entities.User{}, repository.ErrNoUsersFound
	}
	return toModelUser(user), err
}

func toEntityEvent(event dbEvent) entities.Event {
	return entities.Event{
		Id:        event.Id.Hex(),
		Date:      event.Date,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		Content:   event.Content,
		UserId:    event.UserId.Hex(),
	}
}

func (repo *mongoRepo) CreateEvent(params repository.EventOpts) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbId, err := primitive.ObjectIDFromHex(params.UserId)
	if err != nil {
		return "", err
	}
	event := dbEvent{
		Date:      params.Date,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Content:   params.Content,
		UserId:    dbId,
	}
	res, err := coll.InsertOne(ctx, event)
	if err != nil {
		return "", err
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (repo *mongoRepo) GetUserEvents(params struct {
	PeriodStart string
	PeriodEnd   string
	UserId      string
}) ([]entities.Event, error) {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOpts := options.Find()
	findOpts.SetSort(bson.D{{"date", 1}, {"startTime", 1}})

	dbUserId, _ := primitive.ObjectIDFromHex(params.UserId)
	cursor, err := coll.Find(ctx, bson.M{
		"userid": dbUserId,
		"date": bson.D{
			{"$gte", params.PeriodStart},
			{"$lte", params.PeriodEnd},
		},
	}, findOpts)
	var events []entities.Event
	if err != nil {
		return events, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var e dbEvent
		err = cursor.Decode(&e)
		if err != nil {
			return events, err
		}
		events = append(events, toEntityEvent(e))
	}

	return events, err
}

func (repo *mongoRepo) UpdateEvent(params repository.EventOpts) error {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbId, err := primitive.ObjectIDFromHex(params.Id)
	if err != nil {
		return err
	}

	_, err = coll.UpdateByID(ctx, dbId, bson.D{
		{"$set", bson.M{
			"date":      params.Date,
			"startTime": params.StartTime,
			"endTime":   params.EndTime,
			"content":   params.Content,
		}},
	})

	return err
}

func (repo *mongoRepo) DeleteEventById(id string) error {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": dbId})
	return err
}
