.PHONY: build clean install uninstall

export GO111MODULE=on

PREFIX ?= $(HOME)/bin
NAME   ?= ww

build: bin/$(NAME)

clean:
	$(RM) -r bin

install: $(PREFIX)/$(NAME)

uninstall:
	$(RM) $(PREFIX)/$(NAME)

bin/$(NAME): $(shell find . -name *.go)
	go build -o $@ -v github.com/mgirouard/ww/cmd/ww

$(PREFIX)/$(NAME): bin/$(NAME)
	cp $^ $(PREFIX)/$(NAME)
