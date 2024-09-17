package pkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
)

type InvalidURLError struct {
	Cause error
}

func (a InvalidURLError) Error() string {
	return fmt.Errorf("invalid URL address: %w", a.Cause).Error()
}

func Read(htmlPageAddress string) (*html.Node, error) {
	_, err := url.ParseRequestURI(htmlPageAddress)
	if err != nil {
		return nil, InvalidURLError{err}
	}
	response, err := http.Get(htmlPageAddress)
	if err != nil {
		return nil, err
	}
	body := response.Body
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Error(err)
		}
	}(body)
	parsed, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	log.Debugf("HTML Page: %v", parsed)
	return parsed, nil
}
