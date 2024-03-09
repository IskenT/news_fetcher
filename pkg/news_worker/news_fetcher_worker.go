package news_worker

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	models "news_fetcher/internal/domain/entity"
	"news_fetcher/internal/repositories"
	"strconv"
)

type NewsService struct {
	repository *repositories.NewsRepository
}

func NewNewsService(repository *repositories.NewsRepository) *NewsService {
	return &NewsService{
		repository: repository,
	}
}

func (s *NewsService) FetchAndSaveNews(ctx context.Context, fetchUrl string) error {
	// Create a new request with the provided URL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fetchUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing HTTP request: %v", err)
	}

	defer resp.Body.Close()

	var newListInfo models.NewListInformation
	if err := xml.NewDecoder(resp.Body).Decode(&newListInfo); err != nil {
		return fmt.Errorf("error decoding XML data: %v", err)
	}

	// Iterate through each News item
	for _, newsItem := range newListInfo.NewsletterNewsItems {

		// Fetch additional details using NewsArticleID
		additionalDetailsURL := "https://www.htafc.com/api/incrowd/getnewsarticleinformation?id=" + strconv.Itoa(newsItem.NewsArticleID)

		// Create a new request with the additional details URL
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, additionalDetailsURL, nil)
		if err != nil {
			log.Println("Error creating HTTP request for additional details:", err)
			continue // Skip to the next News item
		}

		// Execute the request for additional details
		additionalResp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error fetching additional details:", err)
			continue // Skip to the next News item
		}

		defer additionalResp.Body.Close()

		var additionalInfo models.NewsArticleInformation
		if err := xml.NewDecoder(additionalResp.Body).Decode(&additionalInfo); err != nil {
			log.Println("Error decoding additional details:", err)
			continue // Skip to the next News item
		}

		// Assign additional details to the corresponding News item
		newsItem.Subtitle = additionalInfo.Subtitle
		newsItem.BodyText = additionalInfo.BodyText
		newsItem.GalleryImageURLs = additionalInfo.GalleryImageURLs
		newsItem.VideoURL = additionalInfo.VideoURL

		// Save updated News item into MongoDB
		if err := s.repository.Save(ctx, newsItem); err != nil {
			log.Println("Error saving data to MongoDB:", err)
		}
	}

	return nil
}
