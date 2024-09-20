package service

import (
	"bytes"
	"errors"
	"github.com/h2non/gock"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	// given
	defer gock.Off()
	validAddress := "http://address.com"
	htmlFile, err := os.ReadFile("test_data/html_reader_test.html")
	if err != nil {
		t.Error(err)
	}
	gock.New(validAddress).
		Get("/").
		Reply(200).
		BodyString(string(htmlFile))
	validHtml, err := html.Parse(bytes.NewReader(htmlFile))
	if err != nil {
		t.Error(err)
	}

	invalidAddress := "invalid_address"
	_, invalidURLErr := url.ParseRequestURI(invalidAddress)
	notRealUrl := "http://localhost:01010/test/html_parser"
	_, httpErr := http.Get(notRealUrl)

	type args struct {
		htmlPageAddress string
	}
	cases := []struct {
		name        string
		args        args
		want        *html.Node
		wantErr     bool
		wantErrType error
	}{
		{
			"should be success",
			args{htmlPageAddress: validAddress},
			validHtml,
			false,
			nil,
		},
		{
			"should be failed because of invalid URL",
			args{htmlPageAddress: invalidAddress},
			nil,
			true,
			InvalidURLError{invalidURLErr},
		}, {
			"should be failed because of http call error",
			args{htmlPageAddress: notRealUrl},
			nil,
			true,
			httpErr,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := Read(testCase.args.htmlPageAddress)
			if (err != nil) != testCase.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			if (err != nil) && testCase.wantErr {
				if errors.Is(err, testCase.wantErrType) {
					t.Errorf("Read() should return error = %v, but get %v", testCase.wantErrType, err)
				}
				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Read() got = %v, want %v", got, testCase.want)
			}
		})
	}
}
