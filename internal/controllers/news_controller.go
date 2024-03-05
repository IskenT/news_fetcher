package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	models "news_fetcher/internal/domain/entity"
	"news_fetcher/internal/responses"
	mongoClient "news_fetcher/pkg/client"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var newsCollection *mongo.Collection = mongoClient.GetCollection(mongoClient.DB, "news")

func GetNewsById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := mux.Vars(r)
		newsId := params["newsId"]
		// Преобразуем newsId в int
		newsArticleID, err := strconv.Atoi(newsId)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.NewsResponse{Status: http.StatusBadRequest, Message: "Invalid news ID", Data: nil}
			json.NewEncoder(rw).Encode(response)
			return
		}

		var news models.News

		err = newsCollection.FindOne(ctx, bson.M{"newsarticleid": newsArticleID}).Decode(&news)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				rw.WriteHeader(http.StatusNotFound)
				response := responses.NewsResponse{Status: http.StatusNotFound, Message: "News not found", Data: nil}
				json.NewEncoder(rw).Encode(response)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.NewsResponse{Status: http.StatusInternalServerError, Message: "Error retrieving news", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.NewsResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"news": news}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllNews() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var news []models.News
		defer cancel()

		results, err := newsCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.NewsResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var onePieceOfNews models.News
			if err = results.Decode(&onePieceOfNews); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.NewsResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}

			news = append(news, onePieceOfNews)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.NewsResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": news}}
		json.NewEncoder(rw).Encode(response)
	}
}
