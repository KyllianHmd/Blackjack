##
## EPITECH PROJECT, 2018
## Bollinger
## File description:
## Makfile
##

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=blackjack

all: build

re:

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)

fclean:	clean
	rm -f $(BINARY_NAME)

re:	fclean all

.PHONY: all clean fclean re