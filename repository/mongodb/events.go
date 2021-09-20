package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

func toDtoEvent(event dbEvent) dto.Event {
	return dto.Event{
		Id:        event.Id.Hex(),
		Date:      event.Date,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		Content:   event.Content,
		UserId:    event.UserId.Hex(),
	}
}

func (repo *mongoRepo) CreateEvent(eventData dto.Event) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbId, err := primitive.ObjectIDFromHex(eventData.UserId)
	if err != nil {
		return "", err
	}

	event := dbEvent{
		Date:      eventData.Date,
		StartTime: eventData.StartTime,
		EndTime:   eventData.EndTime,
		Content:   eventData.Content,
		UserId:    dbId,
	}
	res, err := coll.InsertOne(ctx, event)
	if err != nil {
		return "", err
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (repo *mongoRepo) GetEvents(params dto.PeriodParams) ([]dto.Event, error) {
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
	var events []dto.Event
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
		events = append(events, toDtoEvent(e))
	}

	return events, err
}

func (repo *mongoRepo) UpdateEvent(eventData dto.Event) error {
	coll := repo.client.Database(repo.dbName).Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbId, err := primitive.ObjectIDFromHex(eventData.Id)
	if err != nil {
		return err
	}

	_, err = coll.UpdateByID(ctx, dbId, bson.D{
		{"$set", bson.M{
			"date":      eventData.Date,
			"startTime": eventData.StartTime,
			"endTime":   eventData.EndTime,
			"content":   eventData.Content,
		}},
	})

	return err
}

func (repo *mongoRepo) DeleteEvent(id string) error {
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
