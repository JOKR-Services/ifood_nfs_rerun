package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoClient struct {
	mongoUri string
	timeout  time.Duration
}

func NewMongoClient(mongoUri string, timeout time.Duration) *mongoClient {
	return &mongoClient{mongoUri, timeout}
}

func (mc *mongoClient) Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
