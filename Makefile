LAST_COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --tags --abbrev=0)
BUILDSTR := ${VERSION} (\#${LAST_COMMIT} $(shell date -u +"%Y-%m-%dT%H:%M:%S%z"))

BIN := user-service

.PHONY: build
build: $(BIN)

# Build the backend to ./go-user-service.
$(BIN): $(shell find . -type f -name "*.go")
	CGO_ENABLED=0 go build -o ${BIN} -ldflags="-s -w -X 'main.buildString=${BUILDSTR}' -X 'main.versionString=${VERSION}'" cmd/*.go

# Run the backend in dev mode. The frontend assets in dev mode are loaded from disk from frontend/dist/frontend.
.PHONY: run
run:
	CGO_ENABLED=0 go run -ldflags="-s -w -X 'main.buildString=${BUILDSTR}' -X 'main.versionString=${VERSION}' -X 'main.frontendDir=frontend/dist/frontend'" cmd/*.go

# Run Go tests.
.PHONY: test
test:
	go test ./...

# Use goreleaser to do a dry run producing local builds.
.PHONY: release-dry
release-dry:
	goreleaser --parallelism 1 --rm-dist --snapshot --skip-validate --skip-publish

# Use goreleaser to build production releases and publish them.
.PHONY: release
release:
	goreleaser --parallelism 1 --rm-dist --skip-validate
