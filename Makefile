TARGET  = myra-shell
GOPATH ?= /data
GO      = go
GOLINT  = $(GOPATH)/bin/gometalinter

GLIDE_VERSION := $(shell glide --version 2>/dev/null)
DEP_VERSION := $(shell dep version 2>/dev/null)

$(TARGET): clean vendor
	$(GO) build -ldflags="-s -w" -o $@
	upx --brute $@

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

lint: \
	$(GOLINT) \
	$(GOPATH)/bin/goconst \
	$(GOPATH)/bin/ineffassign \
	$(GOPATH)/bin/varcheck \
	$(GOPATH)/bin/structcheck \
	$(GOPATH)/bin/aligncheck \
	$(GOPATH)/bin/gocyclo \
	$(GOPATH)/bin/interfacer \
	$(GOPATH)/bin/gosimple \
	$(GOPATH)/bin/deadcode \
	$(GOPATH)/bin/unconvert \
	$(GOPATH)/bin/staticcheck \
	$(GOPATH)/bin/gotype \
	$(GOPATH)/bin/gas
		$(GOLINT) --deadline 30s $$(find . -type d -not -iwholename '*.git*' | grep -v '/vendor' | tail -n +2)

$(GOPATH)/bin/goconst:
	$(GO) get github.com/jgautheron/goconst/cmd/goconst

$(GOPATH)/bin/ineffassign:
	$(GO) get github.com/gordonklaus/ineffassign

$(GOLINT):
	$(GO) get -u gopkg.in/alecthomas/gometalinter

$(GOPATH)/bin/gotype:
	$(GO) get -u golang.org/x/tools/cmd/gotype

$(GOPATH)/bin/aligncheck:
	$(GO) get github.com/opennota/check/cmd/aligncheck

$(GOPATH)/bin/structcheck:
	$(GO) get github.com/opennota/check/cmd/structcheck

$(GOPATH)/bin/varcheck:
	$(GO) get github.com/opennota/check/cmd/varcheck

$(GOPATH)/bin/gocyclo:
	$(GO) get github.com/fzipp/gocyclo

$(GOPATH)/bin/interfacer:
	$(GO) get github.com/mvdan/interfacer/cmd/interfacer

$(GOPATH)/bin/gosimple:
	$(GO) get honnef.co/go/tools/cmd/gosimple

$(GOPATH)/bin/deadcode:
	$(GO) get github.com/tsenart/deadcode

$(GOPATH)/bin/unconvert:
	$(GO) get github.com/mdempsky/unconvert

$(GOPATH)/bin/staticcheck:
	$(GO) get honnef.co/go/tools/cmd/staticcheck

$(GOPATH)/bin/gas:
	$(GO) get github.com/GoASTScanner/gas

.PHONY:clean test lint
