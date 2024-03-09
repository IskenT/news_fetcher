package repositories

import (
	"context"
	"fmt"
	models "news_fetcher/internal/domain/entity"
	mongoClient "news_fetcher/pkg/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepository struct {
	collection *mongo.Collection
}

func NewNewsRepository(ctx context.Context, client *mongoClient.MongoClient) *NewsRepository {
	return &NewsRepository{
		collection: client.GetCollection("news"),
	}
}

func (r *NewsRepository) Save(ctx context.Context, news models.News) error {
	// Check if the news item already exists in the database
	var existingNews models.News
	err := r.collection.FindOne(ctx, bson.M{"news_article_id": news.NewsArticleID}).Decode(&existingNews)
	if err == nil {
		// If the news item already exists, skip insertion
		fmt.Println("News with ID:", news.NewsArticleID, "already exists. Skipping insertion.")
		return nil
	} else if err != mongo.ErrNoDocuments {
		// If an error other than "no documents found" occurred, return it
		return fmt.Errorf("error finding existing news: %v", err)
	}

	// If the news item doesn't exist, insert it into the database
	_, err = r.collection.InsertOne(ctx, news)
	if err != nil {
		return fmt.Errorf("error inserting news: %v", err)
	}
	fmt.Println("Inserted news with ID:", news.NewsArticleID)

	return nil
}

func (r *NewsRepository) GetById(ctx context.Context, newsId int) (models.News, error) {
	var news models.News
	err := r.collection.FindOne(ctx, bson.M{"news_article_id": newsId}).Decode(&news)
	if err != nil {
		return models.News{}, fmt.Errorf("error finding news by ID: %v", err)
	}
	return news, nil
}

func (r *NewsRepository) GetAll(ctx context.Context) ([]models.News, error) {
	var newsList []models.News

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving news: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		if err := cursor.Decode(&news); err != nil {
			return nil, fmt.Errorf("error decoding news data: %v", err)
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}
