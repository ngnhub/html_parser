package web

import (
	"github.com/bmizerany/pat"
	"github.com/ngnhub/html_scrapper/config"
	"github.com/ngnhub/html_scrapper/service"
	"net/http"
)

func Route() http.Handler {
	context := config.GetAppContext()
	mux := pat.New()

	home := Home{context}
	scrapperService := service.ScrapperService{Searcher: service.DefaultSearcher{}}
	parser := ParserController{service: scrapperService}

	mux.Get("/", &home)
	mux.Post("/", &parser)
	return mux
}
