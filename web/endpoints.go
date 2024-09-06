package web

import (
	"encoding/json"
	"github.com/ngnhub/html_scrapper/config"
	"github.com/ngnhub/html_scrapper/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Home struct {
	context *config.AppContext
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := []byte("Hi!")
	_, err := w.Write(resp)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
	}
}

type ParserController struct {
	context *config.AppContext
	service service.ScrapperService
}

type ParseRequest struct {
	HtmlAddress string   `json:"html_address"`
	Keys        []string `json:"keys"`
}

func (p ParserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := ParseRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		AutoHandle(err, w)
		return
	}
	html, err := service.Read(request.HtmlAddress)
	if err != nil {
		AutoHandle(err, w)
		return
	}
	scrap := p.service.Scrap(request.Keys, html)
	err = json.NewEncoder(w).Encode(scrap)
	if err != nil {
		AutoHandle(err, w) // todo general error handler
		return
	}
	log.Debugf("Parsed request %v", request)
}
