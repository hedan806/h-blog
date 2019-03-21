SHELL = /bin/bash

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=h-blog
BINARY_UNIX=$(BINARY_NAME)-unix

VERSION=0.0.1
DOCKER_NAME=hedan806/h-blog:$(VERSION)

all: build-linux build-docker push-docker

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

build-docker:
	docker build -t $(DOCKER_NAME) .

push-docker:
	docker push $(DOCKER_NAME)

run:
	docker run --rm --net host --name antdu-devops-go $(DOCKER_NAME)
