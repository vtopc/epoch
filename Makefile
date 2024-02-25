.PHONY: help
help: ## This help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run unit tests.
	go test `go list ./... | grep -v '/mocks'` -cover -count=1 -coverprofile=coverage.txt -covermode=count

.PHONY: deps
deps: ## Install dependencies.
	go mod download

# linter:
GOLINT = $(GOPATH)/bin/golangci-lint
$(GOLINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.56.2

.PHONY: lint
lint: $(GOLINT) ## Run linters.
	$(GOLINT) run
