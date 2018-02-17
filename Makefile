NAME     := maguro
VERSION  := v0.3.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
DOCKER_IMAGE_NAME := hinata/maguro-bot
DOCKER_IMAGE_TAG  ?= latest
DOCKER_IMAGE      := $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: clean
clean:
	rm -rf vendor/*

docker-build:
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)_linux
	docker build -t $(DOCKER_IMAGE) .

