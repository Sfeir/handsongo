.DEFAULT_GOAL := dist

OS_MACOS_DIST_NAME=macos
OS_MACOS_NAME=darwin
OS_LINUX_DIST_NAME=linux
OS_LINUX_NAME=linux
OS_ARCH=amd64
GOLANG_VERSION=1.10
GOLANG_DEP_VERSION=v0.4.1
VSCODE_MACOS_VERSION=620882
VSCODE_LINUX_VERSION=620884
VSCODE_GOLANG_PLUGIN_VERSION=0.6.79
JQ_VERSION=1.5
JQ_MACOS_ARCH=osx-amd64
JQ_LINUX_ARCH=linux64
DOCKER_GOLANG_IMAGE_VERSION=1.10-alpine
DOCKER_MONGO_IMAGE_VERSION=3.4

dist: macos linux docker ## Build all distributions
	@rm -rf dist/$(OS_MACOS_DIST_NAME) dist/$(OS_LINUX_DIST_NAME) dist/docker
	@cp INSTALLATION.md dist/

macos: ## Build macos distribution
	@./distribution.sh $(OS_MACOS_DIST_NAME) $(OS_MACOS_NAME) $(OS_ARCH) $(GOLANG_VERSION) $(GOLANG_DEP_VERSION) $(VSCODE_MACOS_VERSION) $(VSCODE_GOLANG_PLUGIN_VERSION) $(JQ_VERSION) $(JQ_MACOS_ARCH)

linux: ## Build linux distribution
	@./distribution.sh $(OS_LINUX_DIST_NAME) $(OS_LINUX_NAME) $(OS_ARCH) $(GOLANG_VERSION) $(GOLANG_DEP_VERSION) $(VSCODE_LINUX_VERSION) $(VSCODE_GOLANG_PLUGIN_VERSION) $(JQ_VERSION) $(JQ_LINUX_ARCH)

docker: ## Build docker distribution
	@./docker.sh $(DOCKER_GOLANG_IMAGE_VERSION) $(DOCKER_MONGO_IMAGE_VERSION)

help: ## Print this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: macos linux dist docker help
