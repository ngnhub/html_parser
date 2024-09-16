package web

import (
	"errors"
	"github.com/ngnhub/html_scrapper/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AutoHandle(err error, w http.ResponseWriter) {
	var invalidUrl service.InvalidURLError
	switch {
	case errors.As(err, &invalidUrl):
		Code400(err, w)
	default:
		Code500(err, w)
	}
}

func Code400(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Invalid request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func Code500(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Error whie handling request %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
