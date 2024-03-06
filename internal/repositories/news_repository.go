package repositories

import (
	"context"
	"fmt"
	models "news_fetcher/internal/domain/entity"
	configs "news_fetcher/pkg/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepository struct {
	collection *mongo.Collection
}

func NewNewsRepository() *NewsRepository {
	return &NewsRepository{
		collection: configs.GetCollection(configs.DB, "news"),
	}
}

func (r *NewsRepository) Save(news models.News) error {
	ctx := context.Background()

	// Check if the news item already exists in the database
	var existingNews models.News
	err := r.collection.FindOne(ctx, bson.M{"newsarticleid": news.NewsArticleID}).Decode(&existingNews)
	if err == nil {
		// If the news item already exists, skip insertion
		fmt.Println("News with ID:", news.NewsArticleID, "already exists. Skipping insertion.")
		return nil
	} else if err != mongo.ErrNoDocuments {
		// If an error other than "no documents found" occurred, return it
		return err
	}

	// If the news item doesn't exist, insert it into the database
	_, err = r.collection.InsertOne(ctx, news)
	if err != nil {
		return err
	}
	fmt.Println("Inserted news with ID:", news.NewsArticleID)

	return nil
}
