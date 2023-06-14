.PHONY: build build-all build-% lint install clean

SHELL := $(shell command -v bash)

ifeq ($(OS),Windows_NT)
    SUFFIX = .exe
    RM = cmd //C rmdir //Q //S
else
    RM = rm -rf
    SUFFIX =
endif

REPO := github.com/BlackHole1/b
VERSION := $(strip $(shell cat ./VERSION))
INPUT := ./cmd/b/
OUTPUT := ./bin

GO ?= go
BUILD_VARS ?= -X $(REPO)/internal/version.Version=$(VERSION)
LDFLAGS ?= -ldflags "-s -w $(BUILD_VARS) --extldflags '-static -fpic'"
GOOS ?= $(strip $(shell go env GOOS))
GOARCH ?= $(strip $(shell go env GOARCH))
GO_CMD ?= CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO)

build:
	$(GO_CMD) build $(LDFLAGS) -o $(OUTPUT)/b-$(GOOS)-$(GOARCH)$(SUFFIX) $(INPUT)

build-%:
	$(eval GOOS := $(firstword $(subst -, ,$*)))
	$(eval GOARCH := $(lastword $(subst -, ,$*)))

	if [[ "$(GOOS)" == "windows" ]]; then \
		$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) SUFFIX=.exe build; \
	else \
  		$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) build; \
  	fi

build-all: clean
	$(MAKE) build-linux-amd64
	$(MAKE) build-linux-arm64
	$(MAKE) build-windows-amd64
	$(MAKE) build-windows-arm64
	$(MAKE) build-darwin-amd64
	$(MAKE) build-darwin-arm64

lint:
	golangci-lint run

install:
	$(eval VERSION := $(addsuffix -dev, $(VERSION)))
	$(GO_CMD) install $(LDFLAGS) $(INPUT)

clean:
	$(RM) ./bin
