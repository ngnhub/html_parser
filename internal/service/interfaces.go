package service

import (
	"github.com/ngnhub/html_scrapper/internal/service/search"
	"golang.org/x/net/html"
)

type ScrapperService interface {
	Scrap(keys []search.Key, source *html.Node) []Found
}
