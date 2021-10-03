package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kuznetsov/event-calendar-backend/dto"
	"github.com/d-kuznetsov/event-calendar-backend/repository"
)

func toDtoUser(user dbUser) dto.User {
	return dto.User{
		Id:       user.Id.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (repo *mongoRepo) CreateUser(userDto dto.User) (string, error) {
	coll := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	user := dbUser{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (repo *mongoRepo) GetUserByEmail(email string) (dto.User, error) {
	coll := repo.client.Database(repo.dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var user dbUser
	err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return dto.User{}, repository.ErrNoUsersFound
	}

	return toDtoUser(user), err
}
