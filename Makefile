BINARY_NAME="imagyn_go"

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

build: ## Build the binary
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server

run: ## Run the server
	$(GORUN) ./cmd/server/main.go

clean: ## Clean the binary
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps: ## Install dependencies
	$(GOGET) -u ./...
	$(GOMOD) tidy

test: ## Run tests
	$(GOTEST) ./... -v

dev: ## Run the server in development mode
	$(GORUN) -tags=dev ./cmd/server/main.go

.PHONY: build run clean deps test dev