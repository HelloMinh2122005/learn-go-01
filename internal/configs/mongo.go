package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	URI        string
	DBName     string
	TimeoutSec int
}

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo(ctx context.Context, cfg MongoConfig) (*Mongo, error) {
	clientOpts := options.Client().ApplyURI(cfg.URI)
	ctx2, cancel := context.WithTimeout(ctx, time.Duration(cfg.TimeoutSec)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx2, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}
	ctx3, cancel2 := context.WithTimeout(ctx, time.Duration(cfg.TimeoutSec)*time.Second)
	defer cancel2()
	if err := client.Ping(ctx3, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping: %w", err)
	}
	db := client.Database(cfg.DBName)
	return &Mongo{
		Client: client,
		DB:     db,
	}, nil
}

func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.DB.Collection(name)
}

func (m *Mongo) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
