package models

type NewListInformation struct {
	ClubName            string `xml:"ClubName"`
	ClubWebsiteURL      string `xml:"ClubWebsiteURL"`
	NewsletterNewsItems []News `xml:"NewsletterNewsItems>NewsletterNewsItem"`
}

type News struct {
	ArticleURL        string `xml:"ArticleURL" json:"article_url"`
	NewsArticleID     int    `xml:"NewsArticleID" json:"news_article_id"`
	PublishDate       string `xml:"PublishDate" json:"publish_date"`
	Taxonomies        string `xml:"Taxonomies" json:"taxonomies"`
	TeaserText        string `xml:"TeaserText" json:"teaser_text"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL" json:"thumbnail_image_url"`
	Title             string `xml:"Title" json:"title"`
	OptaMatchID       string `xml:"OptaMatchId" json:"opta_match_id"`
	LastUpdateDate    string `xml:"LastUpdateDate" json:"last_update_date"`
	IsPublished       bool   `xml:"IsPublished" json:"is_published"`

	// Additional details from NewsArticleInformation
	Subtitle         string   `xml:"Subtitle"`
	BodyText         string   `xml:"BodyText"`
	GalleryImageURLs []string `xml:"GalleryImageURLs"`
	VideoURL         string   `xml:"VideoURL"`
}

type NewsArticleInformation struct {
	Subtitle         string   `xml:"Subtitle"`
	BodyText         string   `xml:"BodyText"`
	GalleryImageURLs []string `xml:"GalleryImageURLs>GalleryImageURL"`
	VideoURL         string   `xml:"VideoURL"`
}
