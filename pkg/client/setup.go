package mongoClient

import (
	"context"
	"fmt"
	config "news_fetcher/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

func NewClient() (*MongoClient, error) {
	cfg := config.Get()
	clientOptions := options.Client().ApplyURI(cfg.MONGODB)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Проверка подключения к базе данных
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB server: %w", err)
	}
	fmt.Println("Connected to MongoDB")

	return &MongoClient{client: client}, nil
}

// GetCollection возвращает коллекцию из базы данных.
func (mc *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return mc.client.Database("golangAPI").Collection(collectionName)
}
