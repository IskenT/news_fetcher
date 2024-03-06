package main

import (
	"fmt"
	"log"
	"net/http"
	"news_fetcher/internal/repositories"
	"news_fetcher/internal/routes"
	mongoClient "news_fetcher/pkg/client"
	"news_fetcher/pkg/news_worker"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

const (
	fetchURL = "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
)

func main() {
	router := mux.NewRouter()

	// Run database
	mongoClient.ConnectDB()

	// Create repository
	newsRepo := repositories.NewNewsRepository()

	// Create service
	newsService := news_worker.NewNewsService(newsRepo)

	c := cron.New()
	// Schedule the fetchAndSaveData function to run every 1 minute
	c.AddFunc("0 */1 * * *", func() {
		fmt.Println("Cron started")
		newsService.FetchAndSaveNews(fetchURL) // Use service method
	})
	c.Start()

	// Routes
	routes.NewsFetcherRoute(router) // Add this

	log.Fatal(http.ListenAndServe(":6000", router))
}
