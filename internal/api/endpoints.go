package api

import (
	"encoding/json"
	"github.com/ngnhub/html_scrapper/internal/config"
	"github.com/ngnhub/html_scrapper/internal/service"
	"github.com/ngnhub/html_scrapper/internal/service/reader"
	"github.com/ngnhub/html_scrapper/internal/service/search"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HomeController struct {
	context *config.Application
}

func (h *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := []byte("Hi!")
	_, err := w.Write(resp)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

type ScrapperController struct {
	service service.ScrapperService
}

type ParseRequest struct {
	HtmlAddress string       `json:"html_address"`
	Keys        []search.Key `json:"keys"`
}

func (p *ScrapperController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := ParseRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		AutoHandle(err, w)
		return
	}
	html, err := reader.Read(request.HtmlAddress)
	if err != nil {
		AutoHandle(err, w)
		return
	}
	scrap := p.service.Scrap(request.Keys, html)
	err = json.NewEncoder(w).Encode(scrap)
	if err != nil {
		AutoHandle(err, w)
		return
	}
	log.Debugf("Parsed request %v", request)
}
