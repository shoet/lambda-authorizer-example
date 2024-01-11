.DEFAULT_GOAL := help

.PHONY: clean
clean: ## clean up
	rm -rf cmd/bin

.PHONY: build-bin
build-bin: ## build go binary
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o cmd/bin/api cmd/api/main.go

.PHONY: deploy
deploy: clean build-bin ## deploy by sls to remote
	sls deploy --verbose

.PHONY: rm
rm: ## remove by sls from remote
	sls remove --verbose

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
