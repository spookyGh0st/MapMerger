# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

GIT_HASH=`git rev-parse HEAD`
GIT_TAG=`git tag --points-at HEAD`
BUILD_TIME=`date +%Y-%m-%dT%T%z`
OUTPUT=build
BUILD_FLAGS=-ldflags="-s -w -X main.sha1ver=$(GIT_HASH) -X main.gitTag=$(GIT_TAG) -X main.buildTime=$(BUILD_TIME)"

BINARY_NAME=MapMerger
BINARY_WIN=$(BINARY_NAME).exe
BINARY_MAC=$(BINARY_NAME)-mac

build: build-win build-linux build-mac

# Build songe-converter
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_WIN) -v .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_NAME) -v .

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_MAC) -v .


clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT)

release:
	@read -p "Enter version: " version; \
	git tag "v$$version" && \
	git push && \
	git push --tags && \
	make clean && \
	make
