install:
	go get golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get github.com/google/wire/cmd/wire
	go install github.com/amacneil/dbmate
	go get -u github.com/smallnest/gen
	go get github.com/golang/mock/mockgen@v1.4.4

lint:
	golint ./...

stop:
	docker-compose -f ./docker/docker-compose.yml down

local: stop
	docker-compose -f ./docker/docker-compose.yml up --build --remove-orphans
