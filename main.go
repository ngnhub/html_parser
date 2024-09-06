package main

import (
	"fmt"
	"github.com/ngnhub/html_scrapper/config"
	"github.com/ngnhub/html_scrapper/web"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var context = config.GetAppContext()
var pops = context.ConfigProperties

func main() {
	context.ConfigLogger()
	properties := pops.ServerProperties
	handler := web.Route()
	server := &http.Server{Addr: fmt.Sprintf(":%v", properties.Port), Handler: handler}
	log.Infof("Server started on %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Can not start server: %v", err)
	}
}
