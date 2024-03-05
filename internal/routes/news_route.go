package routes

import (
	"news_fetcher/internal/controllers"

	"github.com/gorilla/mux"
)

const (
	newsURL    = "/news/{newsId}"
	newsAllURL = "/news"
)

func NewsFetcherRoute(router *mux.Router) {
	router.HandleFunc(newsURL, controllers.GetNewsById()).Methods("GET")
	router.HandleFunc(newsAllURL, controllers.GetAllNews()).Methods("GET")
}
