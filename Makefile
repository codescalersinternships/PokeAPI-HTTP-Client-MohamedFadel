GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pokeapi-cli
DOCKER_IMAGE_NAME=pokeapi-cli

all: deps fmt lint test build docker-build docker-run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET) -v -t -d ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run --rm -it --name "pokeapiclient" --hostname "pokeapiclient" $(DOCKER_IMAGE_NAME) 

.PHONY: all build test clean run deps fmt lint docker-build docker-run
