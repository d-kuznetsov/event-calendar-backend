package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbUser struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type dbEvent struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Date      string             `json:"date"`
	StartTime string             `json:"startTime"`
	EndTime   string             `json:"endTime"`
	Content   string             `json:"content"`
	UserId    primitive.ObjectID `json:"userId"`
}
