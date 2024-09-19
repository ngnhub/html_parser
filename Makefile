test:
	go test ./...

coverage: test
	go test ./... -coverprofile=./coverage/coverage.out
	go tool cover -o=./coverage/coverage.html -html=./coverage/coverage.out

linter :
	golangci-lint run

run:
	go run ./cmd/html_scrapper/main.go