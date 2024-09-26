package api

import (
	"errors"
	"fmt"
	"github.com/ngnhub/html_scrapper/internal/service/reader"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AutoHandle(err error, w http.ResponseWriter) {
	var invalidUrl reader.InvalidURLError
	switch {
	case errors.As(err, &invalidUrl):
		badRequest(err, w)
	default:
		internalError(err, w)
	}
}

func badRequest(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Invalid request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func internalError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Errorf("Error whie handling request %v", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
	}
}
