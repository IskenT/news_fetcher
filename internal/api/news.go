package api

import models "news_fetcher/internal/domain/entity"

// для реквестов и ресопнсов с контроллера
type News struct {
	ArticleURL        string   `json:"article_url"`
	NewsArticleID     int      `json:"news_article_id"`
	PublishDate       string   `json:"publish_date"`
	Taxonomies        string   `json:"taxonomies"`
	TeaserText        string   `json:"teaser_text"`
	ThumbnailImageURL string   `json:"thumbnail_image_url"`
	Title             string   `json:"title"`
	OptaMatchID       string   `json:"opta_match_id"`
	LastUpdateDate    string   `json:"last_update_date"`
	IsPublished       bool     `json:"is_published"`
	Subtitle          string   `json:"subtitle"`
	BodyText          string   `json:"body_text"`
	GalleryImageURLs  []string `json:"gallery_image_urls"`
	VideoURL          string   `json:"video_url"`
}

func GetNewsById(NewsModel *models.News) News {
	return News{
		ArticleURL:        NewsModel.ArticleURL,
		NewsArticleID:     NewsModel.NewsArticleID,
		PublishDate:       NewsModel.PublishDate,
		Taxonomies:        NewsModel.Taxonomies,
		TeaserText:        NewsModel.TeaserText,
		ThumbnailImageURL: NewsModel.ThumbnailImageURL,
		Title:             NewsModel.Title,
		OptaMatchID:       NewsModel.OptaMatchID,
		LastUpdateDate:    NewsModel.LastUpdateDate,
		IsPublished:       NewsModel.IsPublished,
		Subtitle:          NewsModel.Subtitle,
		BodyText:          NewsModel.BodyText,
		GalleryImageURLs:  NewsModel.GalleryImageURLs,
		VideoURL:          NewsModel.VideoURL,
	}
}

func GetNewsList(NewsModels []models.News) []News {
	newsList := make([]News, 0)
	for _, NewsModel := range NewsModels {
		item := News{
			ArticleURL:        NewsModel.ArticleURL,
			NewsArticleID:     NewsModel.NewsArticleID,
			PublishDate:       NewsModel.PublishDate,
			Taxonomies:        NewsModel.Taxonomies,
			TeaserText:        NewsModel.TeaserText,
			ThumbnailImageURL: NewsModel.ThumbnailImageURL,
			Title:             NewsModel.Title,
			OptaMatchID:       NewsModel.OptaMatchID,
			LastUpdateDate:    NewsModel.LastUpdateDate,
			IsPublished:       NewsModel.IsPublished,
			Subtitle:          NewsModel.Subtitle,
			BodyText:          NewsModel.BodyText,
			GalleryImageURLs:  NewsModel.GalleryImageURLs,
			VideoURL:          NewsModel.VideoURL,
		}
		newsList = append(newsList, item)
	}

	return newsList
}
