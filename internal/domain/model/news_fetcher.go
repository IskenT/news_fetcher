package news_fetcher

// NewsFetcher is a object used by fetcher service
type NewListInformationFetcher struct {
	ClubName            string        `xml:"ClubName"`
	ClubWebsiteURL      string        `xml:"ClubWebsiteURL"`
	NewsletterNewsItems []NewsFetcher `xml:"NewsletterNewsItems>NewsletterNewsItem"`
}

type NewsFetcher struct {
	ArticleURL        string   `xml:"ArticleURL"`
	NewsArticleID     int      `xml:"NewsArticleID"`
	PublishDate       string   `xml:"PublishDate"`
	Taxonomies        string   `xml:"Taxonomies"`
	TeaserText        string   `xml:"TeaserText"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
	Title             string   `xml:"Title"`
	OptaMatchID       string   `xml:"OptaMatchId"`
	LastUpdateDate    string   `xml:"LastUpdateDate"`
	IsPublished       bool     `xml:"IsPublished"`
	Subtitle          string   `xml:"Subtitle"`
	BodyText          string   `xml:"BodyText"`
	GalleryImageURLs  []string `xml:"GalleryImageURLs"`
	VideoURL          string   `xml:"VideoURL"`
}

type NewsArticleInformationFetcher struct {
	Subtitle         string   `xml:"Subtitle"`
	BodyText         string   `xml:"BodyText"`
	GalleryImageURLs []string `xml:"GalleryImageURLs>GalleryImageURL"`
	VideoURL         string   `xml:"VideoURL"`
}
