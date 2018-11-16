TARGET  = myra-shell

GOPATH		?= /data
GO      	= go
GOLINT  	= $(GOPATH)/bin/golint
GO_SUBPKGS 	= $(shell $(GO) list ./... | grep -v /vendor/ | sed -e "s!$$($(GO) list)!.!")

GLIDE_VERSION := $(shell glide --version 2>/dev/null)
DEP_VERSION := $(shell dep version 2>/dev/null)

UPX := $(shell upx --version 2>/dev/null)

all: $(TARGET)

$(TARGET): build
ifdef UPX
	upx --brute $@
endif

build: vendor clean
	$(GO) build -ldflags="-s -w" -o $(TARGET) ./main.go

debug: clean vendor
	$(GO) build -gcflags '-N -l'

vendor:
ifdef DEP_VERSION
	dep ensure
else ifdef GLIDE_VERSION
	glide install
else
	go get .
endif

clean:
	rm -f $(TARGET)

test:
	$(GO) test $$($(GO) list ./... | grep -v "/vendor/")

lint: $(GOLINT)
	@for f in $(GO_SUBPKGS) ; do $(GOLINT) $$f ; done

$(GOLINT):
	$(GO) get -u golang.org/x/lint/golint

.PHONY:clean test lint
