package api

import (
	"errors"
	"fmt"
	"github.com/ngnhub/html_scrapper/internal/service/reader"
	"net/http"
	"testing"
)

type MockedResponseWriter struct {
	status int
	body   string
}

func (w *MockedResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *MockedResponseWriter) Write(b []byte) (int, error) {
	w.body = string(b)
	return 0, nil
}

func (w *MockedResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func TestAutoHandle(t *testing.T) {
	type args struct {
		err error
		w   *MockedResponseWriter
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "Should write Bad request in case of Invalid URL error",
			args: args{
				reader.InvalidURLError{Cause: errors.New("some other error")},
				&MockedResponseWriter{},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "invalid URL address: some other error\n",
		},
		{
			name: "Should write Bad request in case of Invalid URL error is not in the top of the chain",
			args: args{
				fmt.Errorf("some top error: %w",
					reader.InvalidURLError{Cause: errors.New("some other error")}),
				&MockedResponseWriter{},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "some top error: invalid URL address: some other error\n",
		},
		{
			name: "Should write Internal Error in case of other errors",
			args: args{
				errors.New("some other error"),
				&MockedResponseWriter{},
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "Internal Server Error: some other error\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := tt.args.w
			AutoHandle(tt.args.err, writer)
			if writer.status != tt.expectedStatusCode {
				t.Errorf("AutoHandle() status = %v, want %v", writer.status, tt.expectedStatusCode)
			}
			if writer.body != tt.expectedBody {
				t.Errorf("AutoHandle() body = %v, want %v", writer.body, tt.expectedBody)
			}
		})
	}
}
