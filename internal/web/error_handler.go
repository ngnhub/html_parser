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
		BadRequest(err, w)
	default:
		InternalError(err, w)
	}
}

func BadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Invalid request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func InternalError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Error whie handling request %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
