BINDIR:=bin
REVISION:=$(shell git rev-parse --short HEAD)

GO_LDFLAGS_VERSION:=-X 'main.Revision=${REVISION}'
GO_LDFLAGS:=$(GO_LDFLAGS_VERSION)
GO_BUILD_OPTION:=-ldflags="-s -w $(GO_LDFLAGS)"

.PHONY: build clean lint test

build: bin/api bin/covapi

$(BINDIR)/api:
	go build -o $@ $(GO_BUILD_OPTION) ./cmd/api/main.go

$(BINDIR)/covapi:
	go build -cover -covermode atomic -o $@ $(GO_BUILD_OPTION) ./cmd/api/main.go

runcovapi: bin/covapi
	mkdir -p covdir
	GOCOVERDIR="$$PWD/covdir" ./bin/covapi

clean:
	go clean -cache -testcache
	rm -f $(BINDIR)/*
	rm -f covdir/*
	rm -f coverage.*
	rm -f treemap.svg

lint:
	go vet ./...

cover: lint
	rm -f coverage.*
	go test ./... -cover -covermode atomic -coverprofile=coverage.out
	go tool cover -html coverage.out -o coverage.html
	go-cover-treemap -coverprofile coverage.out > treemap.svg

coverpkg: lint
	rm -f coverage.*
	go test ./... -cover -coverpkg ./... -covermode atomic -coverprofile=coverage.out
	go tool cover -html coverage.out -o coverage.html
	go-cover-treemap -coverprofile coverage.out > treemap.svg

itcoverage: lint
	rm -f coverage.*
	go test ./... -cover -coverpkg ./... -covermode atomic -test.gocoverdir="$$PWD/covdir"
	go tool covdata textfmt -i=./covdir -o coverage.out
	go tool cover -html coverage.out -o coverage.html
	go-cover-treemap -coverprofile coverage.out > treemap.svg
