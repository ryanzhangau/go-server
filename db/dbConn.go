package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//Connection - dabase connection
func Connection() (client *mongo.Client, err error) {
	if client == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		return client, err
	}
	return client, nil
}
