.PHONY: test
test:
	go test `go list ./... | grep -v '/mocks'` -cover -count=1 -coverprofile=coverage.txt -covermode=count

.PHONY: deps
deps:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod download
