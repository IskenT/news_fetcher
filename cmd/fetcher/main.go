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
	"syscall"
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

	c := cron.New()
	// Schedule the fetchAndSaveData function to run every 1 minute
	c.AddFunc("0 */1 * * *", func() {
		fmt.Println("Cron started")
		newsService.FetchAndSaveNews(ctx, cfg.FETCHUrl) // Use service method
	})
	c.Start()

	// Routes
	routes.NewsFetcherRoute(router)

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	srvErrs := make(chan error, 1)
	go func() {
		srvErrs <- srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdown := stopSrv(srv)

	select {
	case err := <-srvErrs:
		shutdown(err)
	case sig := <-quit:
		shutdown(sig)
	}

	log.Println("Server exiting")
}

func stopSrv(srv *http.Server) func(reason interface{}) {
	return func(reason interface{}) {
		log.Println("Server Shutdown:", reason)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Error Gracefully Shutting Down API:", err)
		}
	}
}
