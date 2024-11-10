package database

import (
	"context"

	"github.com/Lucas-Linhar3s/Rubo/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB is a struct that contains the database connection
type MongoDB struct {
	db *mongo.Database
}

func openMongodb(config *config.Config) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(
		config.Data.DB.User.Driver +
			"://" +
			config.Data.DB.User.Username +
			":" +
			config.Data.DB.User.Password +
			"@" +
			config.Data.DB.User.HostName +
			":" +
			config.Data.DB.User.Port,
	)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		db: client.Database(config.Data.DB.User.Nick),
	}, nil
}

// Close is a function that closes the database connection
func (mg *MongoDB) Close() error {
	return mg.db.Client().Disconnect(context.Background())
}

// GetCollection is a function that returns a collection from the database
func (mg *MongoDB) GetCollection(name string) *mongo.Collection {
	return mg.db.Collection(name)
}
