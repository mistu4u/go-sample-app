install:
	go get github.com/go-chi/chi
	go get github.com/sirupsen/logrus
	go get golang.org/x/net
	go get github.com/stretchr/testify
	go mod download
	go mod verify

lint:install
	golint ./...
run:lint
	go run cmd/main.go
test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
