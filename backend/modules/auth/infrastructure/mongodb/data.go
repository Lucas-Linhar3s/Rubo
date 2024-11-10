package mongodb

import (
	"context"

	"github.com/Lucas-Linhar3s/Rubo/database"
	"github.com/Lucas-Linhar3s/Rubo/modules/auth/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
)

// MGAuth is a struct that represents the mongodb repository of auth
type MGAuth struct {
	Db *database.Database
}

// RegisterUser is a function that registers a new user
func (mg *MGAuth) RegisterUser(model *infrastructure.AuthModel) error {
	collection := mg.Db.Mongo.GetCollection("users")

	_, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}

// VerifyEmail is a function that verifies if an email exists
func (mg *MGAuth) VerifyEmail(email string) (bool, error) {
	collection := mg.Db.Mongo.GetCollection("users")

	cursor, err := collection.Find(context.Background(), bson.M{"email": email})
	if err != nil {
		return false, err
	}

	var exits bool
	if cursor.Next(context.Background()) {
		exits = true
	}

	return exits, nil
}

// LoginWithEmailAndPassword is a function that logs in with email and password
func (mg *MGAuth) LoginWithEmailAndPassword(model *infrastructure.AuthModel) error {
	colletion := mg.Db.Mongo.GetCollection("users")

	cursor, err := colletion.Find(context.Background(), bson.M{"email": model.Email})
	if err != nil {
		return err
	}

	if cursor.Next(context.Background()) {
		if err = cursor.Decode(&model); err != nil {
			return err
		}
	}

	return nil
}

// VerifyRole is a function that verifies the role of a user
func (mg *MGAuth) VerifyRole(model *infrastructure.AuthModel) error {
	collection := mg.Db.Mongo.GetCollection("users")

	cursor, err := collection.Find(context.Background(), bson.M{"email": model.Email})
	if err != nil {
		return err
	}

	if cursor.Next(context.Background()) {
		if err = cursor.Decode(&model); err != nil {
			return err
		}
	}

	return nil
}
