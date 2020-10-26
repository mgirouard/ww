.PHONY: all clean test

export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE=on
export GOFLAGS=-mod=vendor

all: cow

clean:
	$(RM) cow tags

test:
	go test -v .

cow:
	go build -v -o $@ .

tags: $(GOPATH)/bin/gotags
	$< -R -f $@ .

$(GOPATH)/bin/gotags:
	go get -u github.com/jstemmer/gotags
