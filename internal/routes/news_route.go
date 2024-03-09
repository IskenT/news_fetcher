package routes

import (
	"news_fetcher/internal/controllers"
	"news_fetcher/internal/repositories"

	"github.com/gorilla/mux"
)

const (
	newsURL    = "/news/{newsId}"
	newsAllURL = "/news"
)

func NewsFetcherRoute(router *mux.Router, newsRepo *repositories.NewsRepository) {
	router.HandleFunc(newsURL, controllers.GetNewsById(newsRepo)).Methods("GET")
	router.HandleFunc(newsAllURL, controllers.GetAllNews(newsRepo)).Methods("GET")
}
