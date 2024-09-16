package web

import (
	"github.com/bmizerany/pat"
	"github.com/ngnhub/html_scrapper/config"
	"github.com/ngnhub/html_scrapper/service"
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
	scrapperService := service.ScrapperService{Searcher: service.DefaultSearcher{}}
	scrapper := ScrapperController{service: scrapperService}

	mux.Get("/", &home)
	mux.Post("/", &scrapper)
	return mux
}
