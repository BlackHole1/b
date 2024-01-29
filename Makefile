.PHONY: build build-all build-% install clean

REPO := github.com/BlackHole1/b
VERSION := $(strip $(shell cat ./VERSION))
INPUT := ./cmd/b/
OUTPUT := ./out

GO ?= go
BUILD_VARS ?= -X $(REPO)/internal/version.Version=$(VERSION)
LDFLAGS ?= -ldflags "-s -w $(BUILD_VARS) --extldflags '-static -fpic'"
GOOS ?= $(strip $(shell go env GOOS))
GOARCH ?= $(strip $(shell go env GOARCH))
GO_CMD ?= CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO)

build:
	$(GO_CMD) build $(LDFLAGS) -o $(OUTPUT)/b-$(GOOS)-$(GOARCH) $(INPUT)

build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64: build-%:
	$(eval GOOS := $(firstword $(subst -, ,$*)))
	$(eval GOARCH := $(lastword $(subst -, ,$*)))

	$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) build

build-all:
	$(MAKE) build-linux-amd64
	$(MAKE) build-linux-arm64
	$(MAKE) build-darwin-amd64
	$(MAKE) build-darwin-arm64

install:
	$(eval VERSION := $(addsuffix -dev, $(VERSION)))
	$(GO_CMD) install $(LDFLAGS) $(INPUT)

clean:
	$(RM) ./out
