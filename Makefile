.PHONY: build docker-build docker-push release

BIN_DIR := bin
CAMO_CMD := .

build: ## Build camo
	go build -o camo main.go

release: ## Cross compile bins for camo
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/camo-linux-amd64 $(CAMO_CMD)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(BIN_DIR)/camo-linux-arm64 $(CAMO_CMD)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/camo-windows-amd64 $(CAMO_CMD)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o $(BIN_DIR)/camo-windows-arm64 $(CAMO_CMD)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/camo-darwin-amd64 $(CAMO_CMD)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/camo-darwin-arm64 $(CAMO_CMD)

docker-build: build ## Build docker image
	docker build -t ghcr.io/warehouse-13/camo:latest .

docker-push: ## Push docker image
	docker push ghcr.io/warehouse-13/camo:latest

.PHONY: help
help:  ## Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/
ifeq ($(OS),Windows_NT)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  %-40s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
else
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
endif
