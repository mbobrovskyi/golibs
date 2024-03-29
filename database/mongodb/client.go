package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDbClient(ctx context.Context, uri string) (*mongo.Client, error) {
	options := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
