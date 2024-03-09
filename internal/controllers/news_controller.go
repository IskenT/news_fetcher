package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"news_fetcher/internal/api"
	"news_fetcher/internal/repositories"
	"news_fetcher/internal/responses"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNewsById(newsRepo *repositories.NewsRepository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// ваш код
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
		newsItem, err := newsRepo.GetById(r.Context(), newsArticleID)
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

		news := api.GetNewsById(&newsItem)

		rw.WriteHeader(http.StatusOK)
		response := responses.NewsResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"news": news}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllNews(newsRepo *repositories.NewsRepository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		newsList, err := newsRepo.GetAll(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.NewsResponse{Status: http.StatusInternalServerError, Message: "Error retrieving news", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		news := api.GetNewsList(newsList)

		rw.WriteHeader(http.StatusOK)
		response := responses.NewsResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"news": news}}
		json.NewEncoder(rw).Encode(response)
	}
}
