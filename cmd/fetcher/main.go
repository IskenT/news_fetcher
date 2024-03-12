package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	config "news_fetcher/configs"
	"news_fetcher/internal/repositories"
	"news_fetcher/internal/routes"
	mongoClient "news_fetcher/pkg/client"
	"news_fetcher/pkg/news_worker"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {
	ctx := context.Background()

	cfg := config.Get()

	router := mux.NewRouter()

	// Run database
	client, err := mongoClient.NewClient()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Create repository
	newsRepo := repositories.NewNewsRepository(ctx, client)

	// Create service
	newsService := news_worker.NewNewsService(newsRepo)

	crn := cron.New()
	// Schedule the fetchAndSaveData function to run every 1 minute
	crn.AddFunc("0 */1 * * *", func() {
		fmt.Println("Cron started")
		newsService.FetchAndSaveNews(ctx, cfg.FETCHUrl) // Use service method
	})
	crn.Start()

	// Routes
	routes.NewsFetcherRoute(router, newsRepo)

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shut down gracefuly!")
}
