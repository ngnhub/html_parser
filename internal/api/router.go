package api

import (
	"github.com/bmizerany/pat"
	"github.com/ngnhub/html_scrapper/internal/config"
	"github.com/ngnhub/html_scrapper/internal/service"
	"github.com/ngnhub/html_scrapper/internal/service/search/default"
	"net/http"
)

type Router struct {
	app *config.Application
}

func NewRouter(app *config.Application) *Router {
	return &Router{app: app}
}

func (Router) Route() http.Handler {
	context := config.CreateApplication()
	mux := pat.New()

	home := HomeController{context}
	scrapperService := service.PatternDetectScrapperService{Searcher: defaultsearcher.DefaultSearcher{}}
	scrapper := ScrapperController{service: scrapperService}

	mux.Get("/", &home)
	mux.Post("/", &scrapper)
	return mux
}
