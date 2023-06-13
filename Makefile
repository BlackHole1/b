.PHONY: clean

ifeq ($(OS),Windows_NT)
    RM = cmd //C rmdir //Q //S
else
    RM = rm -rf
endif

ifeq ($(OS),Windows_NT)
    SUFFIX = .exe
else
    SUFFIX =
endif

INPUT = ./cmd/b/
OUTPUT = ./bin

GO ?= go
LDFLAGS = -ldflags "-s -w --extldflags '-static -fpic'"
GOOS = $(strip $(shell go env GOOS))
GOARCH = $(strip $(shell go env GOARCH))
GOCMD = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO)

build:
	$(GOCMD) build $(LDFLAGS) -o $(OUTPUT)/b-$(GOOS)-$(GOARCH)$(SUFFIX) $(INPUT)

build-%:
	$(eval GOOS := $(firstword $(subst -, ,$*)))
	$(eval GOARCH := $(lastword $(subst -, ,$*)))

	if [[ "$(GOOS)" == "windows" ]]; then \
		$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) SUFFIX=.exe build; \
	else \
  		$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) build; \
  	fi

build-all:
	$(MAKE) clean
	$(MAKE) build-linux-amd64
	$(MAKE) build-linux-arm64
	$(MAKE) build-windows-amd64
	$(MAKE) build-windows-arm64
	$(MAKE) build-darwin-amd64
	$(MAKE) build-darwin-arm64

clean:
	$(RM) ./bin
