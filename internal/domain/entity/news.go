package models

import (
	models "news_fetcher/internal/domain/model"
)

type News struct {
	ArticleURL        string   `bson:"article_url, omitempty"`
	NewsArticleID     int      `bson:"news_article_id, omitempty"`
	PublishDate       string   `bson:"publish_date, omitempty"`
	Taxonomies        string   `bson:"taxonomies, omitempty"`
	TeaserText        string   `bson:"teaser_text, omitempty"`
	ThumbnailImageURL string   `bson:"thumbnail_image_url, omitempty"`
	Title             string   `bson:"title, omitempty"`
	OptaMatchID       string   `bson:"opta_match_id, omitempty"`
	LastUpdateDate    string   `bson:"last_update_date, omitempty"`
	IsPublished       bool     `bson:"is_published, omitempty"`
	Subtitle          string   `bson:"subtitle, omitempty"`
	BodyText          string   `bson:"body_text, omitempty"`
	GalleryImageURLs  []string `bson:"gallery_image_urls, omitempty"`
	VideoURL          string   `bson:"video_url, omitempty"`
}

func FetchedNewsModelToEntity(NewsModel *models.NewsFetcher) News {
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
