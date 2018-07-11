# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=api
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -v -tags=jsoniter
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -v
	./$(BINARY_NAME)
start:
	fresh
swagger:
	rm -rf docs
	swag init -g server.go
sdk:
	swagger-codegen generate -i docs/swagger/swagger.json -l typescript-fetch -o sdk

