package main

import (
	"fmt"
	"github.com/ngnhub/html_scrapper/internal/api"
	"github.com/ngnhub/html_scrapper/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	app := config.CreateApplication()
	properties := app.ConfigProperties
	router := api.NewRouter(app)
	handler := router.Route()
	server := &http.Server{Addr: fmt.Sprintf(":%v", properties.ServerProperties.Port), Handler: handler}
	log.Infof("Server started on %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Can not start server: %v", err)
	}
}
