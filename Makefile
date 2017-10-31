# Makefile for handsongo : Hands On Go

.DEFAULT_GOAL := all

# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# go env vars
GO=$(firstword $(subst :, ,$(GOPATH)))
# list of pkgs for the project without vendor
PKGS=$(shell go list ./... | grep -v /vendor/)
DOCKER_IP=$(shell if [ -z "$(DOCKER_MACHINE_NAME)" ]; then echo 'localhost'; else docker-machine ip $(DOCKER_MACHINE_NAME); fi)
export GO15VENDOREXPERIMENT=1


# -----------------------------------------------------------------
#        Version
# -----------------------------------------------------------------

# version
VERSION=0.0.5
BUILDDATE=$(shell date -u '+%s')
BUILDHASH=$(shell git rev-parse --short HEAD)
VERSION_FLAG=-ldflags "-X main.Version=$(VERSION) -X main.GitHash=$(BUILDHASH) -X main.BuildStmp=$(BUILDDATE)"

# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

all: clean build ## Clean and build the project

clean: ## Clean the project
	@go clean
	@rm -Rf .tmp .DS_Store *.log *.out *.mem *.test build/

build: format ## Build all libraries and binaries
	@go build -v $(VERSION_FLAG) -o $(GO)/bin/handsongo handsongo.go

format: ## Format all packages
	@go fmt $(PKGS)

teardownTest: ## Tear down mongodb for integration tests
	@$(shell docker kill handsongo-mongo-test 2&>/dev/null 1&>/dev/null)
	@$(shell docker rm handsongo-mongo-test 2&>/dev/null 1&>/dev/null)

setupTest: teardownTest ## Start mongodb for integration tests
	@docker run -d --name handsongo-mongo-test -p "27017:27017" mongo:3.4

test: setupTest ## Start tests with a mongodb docker image
	@export MONGODB_SRV=mongodb://$(DOCKER_IP)/spirits; go test -v $(PKGS); make teardownTest

bench: ## Start benchmark
	@go test -v -run TestSpiritHandlerGet -bench=. -memprofile=prof.mem github.com/Sfeir/handsongo/web

benchTool: bench ## Start benchmark tool
	@echo "### TIP : type 'top 5' and 'list the first item'"
	@go tool pprof --alloc_space web.test prof.mem

lint: ## Lint all packages
	@golint dao/...
	@golint model/...
	@golint web/...
	@golint utils/...
	@golint ./.
	@go vet $(PKGS)

start: ## Start the program
	@$(GO)/bin/handsongo -port 8020 -logl debug -logf text -statd 15s -db mongodb://$(DOCKER_IP)/spirits

stop: ## Stop the program
	@killall handsongo

# -----------------------------------------------------------------
#        Docker targets
# -----------------------------------------------------------------

dockerBuild: ## Build a docker image of the program
	docker build -t sfeir/handsongo:latest .

dockerBuildMulti: ## Build a docker multistep image of the program
	docker build -f Dockerfile.multi -t sfeir/handsongo:latest .

dockerClean: ## Remove the docker image of the program
	docker rmi -f sfeir/handsongo:latest

dockerUp: ## Start the program with its mongodb
	docker-compose up -d

dockerDown: ## Stop the program and the mongodb and remove the containers
	docker-compose down

dockerBuildUp: dockerDown dockerBuild dockerUp ## Stop, build and launch the docker images of the program

dockerBuildUpMulti: dockerDown dockerBuildMulti dockerUp ## Stop, build multi step and launch the docker images of the program

dockerWatch: ## Watch the status of the docker container
	@watch -n1 'docker ps | grep handsongo'

dockerLogs: ## Print the logs of the container
	docker-compose logs -f

help: ## Print this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: all test clean teardownTest setupTest
