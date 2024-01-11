.DEFAULT_GOAL := help

.PHONY: clean
clean: ## clean up
	rm -rf .bin

.PHONY: build
build: ## build go binary
	sls package

.PHONY: deploy
deploy: clean build ## deploy by sls to remote
	sls deploy --verbose

.PHONY: rm
rm: ## remove by sls from remote
	sls remove --verbose

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
