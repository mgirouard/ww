.PHONY: install uninstall

PREFIX ?= $(HOME)/bin
NAME   ?= ww

install: $(PREFIX)/$(NAME)

uninstall:
	$(RM) $(PREFIX)/$(NAME)

$(PREFIX)/$(NAME):
	cp ww.sh $(PREFIX)/$(NAME)
