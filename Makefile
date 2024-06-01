DCOMPOSE=docker-compose

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

deps: 
	$(GOGET) -v ./...
	$(GOMOD) tidy

up: 
	$(DCOMPOSE) up --build

down: 
	$(DCOMPOSE) down