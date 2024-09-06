package service

import "golang.org/x/net/html"

type ScrapperService struct {
	Searcher Searcher
}

type Scrapper interface {
	Scrap(keys []string, source *html.Node) *[]Found
}
