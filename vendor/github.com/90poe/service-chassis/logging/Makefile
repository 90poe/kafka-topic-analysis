export GO111MODULE=on

.PHONY: ci
ci: build test lint

.PHONY: all
all: deps gen build test lint

.PHONY: deps
deps:
	go mod download

.PHONY: build
build:
	go build ./...

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: lint
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
	golangci-lint run -v -c=../.golangci.yml ./...

.PHONY: gen
gen:
	go install github.com/golang/mock/mockgen
	@mkdir -p mocklogging
	mockgen -source=logging.go -destination=mocklogging/mock.go -package mocklogging
