package news_worker

import (
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

func (s *NewsService) FetchAndSaveNews(fetchUrl string) {
	// Fetch XML data and parse it
	resp, err := http.Get(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var newListInfo models.NewListInformation
	if err := xml.NewDecoder(resp.Body).Decode(&newListInfo); err != nil {
		log.Fatal(err)
	}

	// Iterate through each News item
	for _, newsItem := range newListInfo.NewsletterNewsItems {

		// Fetch additional details using NewsArticleID
		additionalDetailsURL := "https://www.htafc.com/api/incrowd/getnewsarticleinformation?id=" + strconv.Itoa(newsItem.NewsArticleID)

		additionalResp, err := http.Get(additionalDetailsURL)
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

		fmt.Println(additionalInfo.BodyText)
		// Assign additional details to the corresponding News item
		newsItem.Subtitle = additionalInfo.Subtitle
		newsItem.BodyText = additionalInfo.BodyText
		newsItem.GalleryImageURLs = additionalInfo.GalleryImageURLs
		newsItem.VideoURL = additionalInfo.VideoURL

		// Save updated News item into MongoDB
		if err := s.repository.Save(newsItem); err != nil {
			log.Println("Error saving data to MongoDB:", err)
		}
	}
}
