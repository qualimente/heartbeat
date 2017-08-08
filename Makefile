NAME=heartbeat
ORG=qualimente
PACKAGE_NAME=github.com/$(ORG)/$(NAME)
TAG=$$(git describe --abbrev=0 --tags)

LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildVersion=$(shell git describe --abbrev=0 --tags)"
LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildSHA=$(shell git rev-parse HEAD)"
# Strip debug information
LDFLAGS += -s

ifeq ($(OS),Windows_NT)
	suffix := .exe
endif

all: build test

$(GOPATH)/bin/glide$(suffix):
	go get github.com/Masterminds/glide

$(GOPATH)/bin/$(NAME)$(suffix):
	go get github.com/$(ORG)/$(NAME)

glide.lock: glide.yaml $(GOPATH)/bin/glide$(suffix)
	glide update
	@touch $@

vendor: glide.lock
	glide install
	@touch $@

releases:
	mkdir -p releases

bin/linux/amd64:
	mkdir -p bin/linux/amd64

bin/darwin/amd64:
	mkdir -p bin/darwin/amd64

build: darwin linux

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

darwin: vendor releases bin/darwin/amd64
	env GOOS=darwin GOAARCH=amd64 go build -ldflags '$(LDFLAGS)' -v -o $(CURDIR)/bin/darwin/amd64/$(NAME)
	tar -cvzf releases/$(NAME)-darwin-amd64.tar.gz bin/darwin/amd64/$(NAME)

linux: vendor releases bin/linux/amd64
	env GOOS=linux GOAARCH=amd64 go build -ldflags '$(LDFLAGS)' -v -o $(CURDIR)/bin/linux/amd64/$(NAME)
	tar -cvzf releases/$(NAME)-linux-amd64.tar.gz bin/linux/amd64/$(NAME)

clean:
	rm -rf releases bin

